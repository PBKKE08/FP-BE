package repository

import (
	"context"
	"github.com/PBKKE08/FP-BE/core/model/partner"
	"github.com/jmoiron/sqlx"
)

type PartnerRepository struct {
	db *sqlx.DB
}

func NewPartnerRepository(db *sqlx.DB) *PartnerRepository {
	return &PartnerRepository{db}
}

func (r *PartnerRepository) ByID(ctx context.Context, id partner.ID) (partner.Partner, error) {
	q := `
		SELECT
			  p.id,
			  p.name,
			  p.email,
			  p.telephone,
			  p.gender,
			  p.price,
			  c.id,
			  c.name
		FROM
			  partners p
			  JOIN cities c ON p.city_id = c.id
		WHERE p.id = ?;
	`

	var result partner.Partner

	err := r.db.QueryRowx(q, id.String()).Scan(
		&result.ID,
		&result.Nama,
		&result.Email,
		&result.NomorTelepon,
		&result.JenisKelamin,
		&result.Harga,
		&result.Kota.ID,
		&result.Kota.Nama,
	)

	if err != nil {
		return result, err
	}

	return result, nil
}

func (r *PartnerRepository) Save(ctx context.Context, p partner.Partner) error {
	qr :=
		`
		INSERT INTO partners (id, name, email, telephone, gender, category_id, price, city_id, description, is_approved)
		VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?);
		`

	_, err := r.db.ExecContext(ctx, qr, p.ID, p.Nama, p.Email, p.NomorTelepon, p.JenisKelamin, p.Kategori.ID.String(), false)
	if err != nil {
		return err
	}

	return nil
}
