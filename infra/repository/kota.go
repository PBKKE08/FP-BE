package repository

import (
	"context"
	"github.com/PBKKE08/FP-BE/core/model/kota"
	"github.com/jmoiron/sqlx"
)

type Kota struct {
	db *sqlx.DB
}

func NewKota(db *sqlx.DB) *Kota {
	return &Kota{db: db}
}

func (k *Kota) ByID(ctx context.Context, id kota.ID) (kota.Kota, error) {
	q := `
		SELECT
			  id,
			  name
		FROM
			  cities
		WHERE id = ?;
	`
	var result kota.Kota

	err := k.db.QueryRowx(q, id.String()).Scan(
		&result.ID,
		&result.Nama,
	)

	if err != nil {
		return result, err
	}

	return result, nil
}
