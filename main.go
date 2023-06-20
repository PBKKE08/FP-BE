package main

import (
	"context"
	command "github.com/PBKKE08/FP-BE/api/command/beri_review"
	"github.com/PBKKE08/FP-BE/api/handler"
	"github.com/PBKKE08/FP-BE/api/usecase"
	"github.com/PBKKE08/FP-BE/infra/query"
	"github.com/PBKKE08/FP-BE/infra/repository"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/PBKKE08/FP-BE/infra/instance"
	"github.com/PBKKE08/FP-BE/pkg"
	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

type Config struct {
	DbURL       string `env:"DB_URL" default:"admin:password@tcp(localhost:3306)/socium_rentalis"`
	ServerPort  string `env:"PORT" default:"7777"`
	Environment string `env:"ENVIRONMENT" default:"DEV"`
}

const maxShutdownTimeout = 1 * time.Minute

var config Config

func init() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	zerolog.SetGlobalLevel(zerolog.InfoLevel)

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

	queryInstance := query.NewQuery(db)
	partnerRepo := repository.NewPartnerRepository(db)
	penggunaRepo := repository.NewPenggunaRepository(db)
	reviewRepo := repository.NewReviewRepository(db)

	beriReviewCmd := command.BeriReview{
		PenggunaRepo: penggunaRepo,
		PartnerRepo:  partnerRepo,
		ReviewRepo:   reviewRepo,
	}

	penggunaUsecase := usecase.NewPenggunaUsecase(queryInstance, &beriReviewCmd)
	penggunaHandler := handler.NewPenggunaHandler(penggunaUsecase)

	server := echo.New()
	penggunaHandler.Load(server)

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
