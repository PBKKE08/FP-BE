package instance

import (
	"context"
	"database/sql"
	_ "github.com/lib/pq"
	"github.com/rs/zerolog/log"
	"time"
)

const connectDBTimeout = 30 * time.Second

type CloseFunc func() error

func NewPostgreSQL(uri string) (*sql.DB, CloseFunc) {
	ctx, cancel := context.WithTimeout(context.Background(), connectDBTimeout)
	defer cancel()

	db, err := sql.Open("postgres", uri)
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
