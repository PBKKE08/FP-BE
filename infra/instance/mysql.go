package instance

import (
	"context"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/rs/zerolog/log"
)

const connectDBTimeout = 30 * time.Second

type CloseFunc func() error

func NewMySQL(uri string) (*sqlx.DB, CloseFunc) {
	ctx, cancel := context.WithTimeout(context.Background(), connectDBTimeout)
	defer cancel()

	db, err := sqlx.Connect("mysql", uri)
	if err != nil {
		log.Fatal().Msgf("can't open new sql conn: %v", err)
	}

	err = db.PingContext(ctx)
	if err != nil {
		log.Fatal().Msgf("can't ping sql conn: %v", err)
	}

	return db, func() error {
		return db.Close()
	}
}
