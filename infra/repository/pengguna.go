package repository

import (
	"context"
	"github.com/PBKKE08/FP-BE/core/model/pengguna"
	"github.com/jmoiron/sqlx"
)

type PenggunaRepository struct {
	db *sqlx.DB
}

func NewPenggunaRepository(db *sqlx.DB) *PenggunaRepository {
	return &PenggunaRepository{db}
}

func (r *PenggunaRepository) ByID(ctx context.Context, id pengguna.ID) (pengguna.Pengguna, error) {
	q := `
		SELECT
			  p.id,
			  p.name,
			  p.email,
			  p.telephone,
			  p.gender,
			  c.id,
			  c.name
		FROM
			  users p
			  JOIN cities c ON p.city_id = c.id
		WHERE p.id = ?;
	`
	var result pengguna.Pengguna

	err := r.db.QueryRowx(q, id.String()).Scan(
		&result.ID,
		&result.Nama,
		&result.Email,
		&result.NomorTelepon,
		&result.JenisKelamin,
		&result.Kota.ID,
		&result.Kota.Nama,
	)

	if err != nil {
		return result, err
	}

	return result, nil
}

func (r *PenggunaRepository) Save(ctx context.Context, user pengguna.Pengguna) error {
	qr :=
		`
		INSERT INTO users (id, name, email, telephone, gender, city_id)
		VALUES (?, ?, ?, ?, ?, ?);
		`

	_, err := r.db.ExecContext(ctx, qr, user.ID, user.Nama, user.Email, user.NomorTelepon, user.JenisKelamin, user.Kota.ID)
	if err != nil {
		return err
	}

	return nil
}
