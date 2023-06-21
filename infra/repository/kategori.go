package repository

import (
	"context"
	"github.com/PBKKE08/FP-BE/core/model/kategori"
	"github.com/jmoiron/sqlx"
)

type Kategori struct {
	db *sqlx.DB
}

func NewKategori(db *sqlx.DB) *Kota {
	return &Kota{db: db}
}

func (k *Kategori) ByID(ctx context.Context, id kategori.ID) (kategori.Kategori, error) {
	q := `
		SELECT
			  id,
			  name
		FROM
			  categories
		WHERE id = ?;
	`
	var result kategori.Kategori

	err := k.db.QueryRowx(q, id.String()).Scan(
		&result.ID,
		&result.Name,
	)

	if err != nil {
		return result, err
	}

	return result, nil
}
