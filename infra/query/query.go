package query

import (
	"context"
	"github.com/PBKKE08/FP-BE/api/query"
	"github.com/jmoiron/sqlx"
)

type Query struct {
	db *sqlx.DB
}

func NewQuery(db *sqlx.DB) *Query {
	return &Query{db: db}
}

func (q *Query) By(ctx context.Context, daerah string, jenisKelamin string) []query.CariPasangan {
	var results = make([]query.CariPasangan, 0)

	switch {
	case daerah != "" && jenisKelamin != "":
		qr := `
		SELECT
		  p.id,
		  p.name,
		  p.price,
		  c.name as c_name,
		  r.rating
		FROM
		  partners p
		  JOIN cities c ON p.city_id = c.id
		  JOIN reviews r ON p.id = r.partner_id
		WHERE
		  c.name = ?
		  AND p.gender = ?;
		`

		q.db.GetContext(ctx, &results, qr, daerah, jenisKelamin)

		return results

	case daerah != "":
		qr := `
		SELECT
		  p.id,
		  p.name,
		  p.price,
		  c.name as c_name,
		  r.rating
		FROM
		  partners p
		  JOIN cities c ON p.city_id = c.id
		  JOIN reviews r ON p.id = r.partner_id
		WHERE
		  c.name = ?;
		`

		q.db.GetContext(ctx, &results, qr, daerah)
		return results

	case jenisKelamin != "":
		qr := `
		SELECT
		  p.id,
		  p.name,
		  p.price,
		  c.name as c_name,
		  r.rating
		FROM
		  partners p
		  JOIN cities c ON p.city_id = c.id
		  JOIN reviews r ON p.id = r.partner_id
		WHERE
		  p.gender = ?;
		`

		q.db.GetContext(ctx, &results, qr, jenisKelamin)
		return results

	default:
		qr := `
		SELECT
		  p.id,
		  p.name,
		  p.price,
		  c.name as c_name,
		  r.rating
		FROM
		  partners p
		  JOIN cities c ON p.city_id = c.id
		  JOIN reviews r ON p.id = r.partner_id;
		`

		q.db.GetContext(ctx, &results, qr)

		return results
	}
}
