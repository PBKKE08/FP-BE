package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/PBKKE08/FP-BE/api/command/beri_review"
	"github.com/PBKKE08/FP-BE/api/command/buat_booking"
	"github.com/PBKKE08/FP-BE/api/command/buat_partner"
	"github.com/PBKKE08/FP-BE/api/command/buat_user"
	"github.com/PBKKE08/FP-BE/api/command/terima_partner"
	"github.com/PBKKE08/FP-BE/api/command/tolak_partner"
	"github.com/PBKKE08/FP-BE/api/handler"
	"github.com/PBKKE08/FP-BE/api/usecase"
	"github.com/PBKKE08/FP-BE/infra/authentication"
	"github.com/PBKKE08/FP-BE/infra/mailer"
	"github.com/PBKKE08/FP-BE/infra/query"
	"github.com/PBKKE08/FP-BE/infra/repository"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4/middleware"
	"google.golang.org/api/option"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	firebase "firebase.google.com/go/v4"
	"github.com/PBKKE08/FP-BE/infra/instance"
	"github.com/PBKKE08/FP-BE/pkg"
	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

type Config struct {
	DbURL        string `env:"DB_URL" default:"admin:password@tcp(localhost:3306)/socium_rentalis"`
	ServerPort   string `env:"PORT" default:"7777"`
	Environment  string `env:"ENVIRONMENT" default:"DEV"`
	MailHost     string `env:"MAIL_HOST" default:"localhost"`
	MailPort     int    `env:"MAIL_PORT" default:"1025"`
	MailUsername string `env:"MAIL_USERNAME" default:"debuggerMail"`
	MailPassword string `env:"MAIL_PASSWORD" default:""`
	MailEmail    string `env:"MAIL_EMAIL" default:"info@company.com"`
}

const maxShutdownTimeout = 1 * time.Minute

var config Config

func init() {
	var isProd bool

	flag.BoolVar(&isProd, "dev", true, "Specify if the project is in production mode")

	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	zerolog.SetGlobalLevel(zerolog.InfoLevel)

	if !isProd {
		godotenv.Load()
	}

	err := pkg.FillEnv(&config)
	if err != nil {
		log.Fatal().Msgf("can't parse env => %v", err)
	}

	if config.Environment == "DEV" {
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
		log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	}
}

func main() {
	db, closeDB := instance.NewMySQL(config.DbURL)
	defer func() {
		if err := closeDB(); err != nil {
			log.Error().Err(err)
		}
	}()

	mailer.SetHost(config.MailHost)
	mailer.SetUsername(config.MailUsername)
	mailer.SetPassword(config.MailPassword)
	mailer.SetEmailServerURI(fmt.Sprintf("%s:%d", config.MailHost, config.MailPort))

	opt := option.WithCredentialsFile("sak.json")
	app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		log.Fatal().Msgf("Error initializing Firebase app: %v\n", err)
	}

	partnerRepo := repository.NewPartnerRepository(db)
	penggunaRepo := repository.NewPenggunaRepository(db)
	reviewRepo := repository.NewReviewRepository(db)
	kotaRepo := repository.NewKota(db)
	txRepo := repository.NewTransactionRepository(db)
	orderRepo := repository.NewOrderRepo(db)
	kategoriRepo := repository.NewKategori(db)

	queryInstance := query.NewQuery(db)
	mailer := mailer.Mailer(mailer.SendEmail)
	jwtProvider := authentication.JWTProvider(authentication.GenerateToken)

	authInstance, err := authentication.NewFirebaseAuth(app)
	if err != nil {
		log.Fatal().Msgf("Err creating auth instance: %v\n", err)
	}

	beriReviewCmd := beri_review.BeriReview{
		PenggunaRepo: penggunaRepo,
		PartnerRepo:  partnerRepo,
		ReviewRepo:   reviewRepo,
	}

	buatUserCmd := buat_user.BuatUser{
		PenggunaRepo: penggunaRepo,
		KotaRepo:     kotaRepo,
	}

	buatPartnerCmd := buat_partner.BuatPartner{
		PartnerRepo:  partnerRepo,
		KotaRepo:     kotaRepo,
		KategoriRepo: kategoriRepo,
	}

	buatBookingCmd := buat_booking.BuatBooking{
		TransactionRepo: txRepo,
		OrderRepo:       orderRepo,
		PenggunaRepo:    penggunaRepo,
		PartnerRepo:     partnerRepo,
	}

	terimaPartnerCmd := terima_partner.TerimaPartner{PartnerRepo: partnerRepo}
	tolakPartnerCmd := tolak_partner.TolakPartner{PartnerRepo: partnerRepo}

	publicUsecase := usecase.NewPublicUsecase(queryInstance)
	publicHandler := handler.NewPublicHandler(publicUsecase)

	penggunaUsecase := usecase.NewPenggunaUsecase(queryInstance, &beriReviewCmd, queryInstance, queryInstance)
	penggunaHandler := handler.NewPenggunaHandler(penggunaUsecase)

	partnerUsecase := usecase.NewPartnerUsecase(queryInstance)
	partnerHandler := handler.NewPartnerHandler(partnerUsecase)

	bookingUsecase := usecase.NewBookingUsecase(&buatBookingCmd)
	bookingHandler := handler.NewBookingHandler(bookingUsecase)

	authUsecase := usecase.NewAuthUsecase(&buatUserCmd, authInstance, queryInstance, mailer, jwtProvider, &buatPartnerCmd, queryInstance)
	authHandler := handler.NewAuthHandler(authUsecase)

	adminUsecase := usecase.NewAdminUsecase(queryInstance, &tolakPartnerCmd, &terimaPartnerCmd, queryInstance, authInstance)
	adminHandler := handler.NewAdminHandler(adminUsecase)

	server := echo.New()
	server.Use(middleware.CORS())
	server.Use(middleware.Recover())
	server.Use(middleware.Logger())

	bookingHandler.Load(server)
	penggunaHandler.Load(server)
	authHandler.Load(server)
	publicHandler.Load(server)
	partnerHandler.Load(server)
	adminHandler.Load(server)

	go func() {
		if err := server.Start(":" + config.ServerPort); err != nil && err != http.ErrServerClosed {
			log.Fatal().Msgf("can't start server: %v", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	<-quit

	ctx, cancel := context.WithTimeout(context.Background(), maxShutdownTimeout)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		log.Fatal().Msgf("can't shutdown server: %v", err)
	}

	log.Info().Msg("Server closed")
}
