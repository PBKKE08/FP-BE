package main

import (
	"context"
	"github.com/PBKKE08/FP-BE/echo-rest/routes"
	"github.com/PBKKE08/FP-BE/infra/instance"
	"github.com/PBKKE08/FP-BE/pkg"
	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

type Config struct {
	DbURL       string `env:"DB_URL" default:"postgresql://postgres:postgres@localhost:5678/socium_rentalis?sslmode=disable"`
	RedisURL    string `env:"REDIS_URL" default:"localhost:6379"`
	ServerPort  string `env:"PORT" default:"6666"`
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
	db, closeDB := instance.NewPostgreSQL(config.DbURL)
	_, closeRedis := instance.NewRedis(config.RedisURL)
	defer func() {
		if err := closeDB(); err != nil {
			log.Error().Err(err)
		}

		if err := closeRedis(); err != nil {
			log.Error().Err(err)
		}
	}()

	server := echo.New()
	routes.GetRoutes(db, server)

	go func() {
		if err := server.Start(":" + config.ServerPort); err != nil && err != http.ErrServerClosed {
			log.Fatal().Msgf("can't start server: %v", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGINT, syscall.SIGKILL, syscall.SIGTERM)

	<-quit

	ctx, cancel := context.WithTimeout(context.Background(), maxShutdownTimeout)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		log.Fatal().Msgf("can't shutdown server: %v", err)
	}

	log.Info().Msg("Server closed")
}
