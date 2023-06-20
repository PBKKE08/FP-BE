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
	var results []query.CariPasangan

	switch {
	case daerah != "" && jenisKelamin != "":
		qr := `
		SELECT
		  p.id,
		  p.name,
		  p.price,
		  c.name as c_name,
		  COALESCE(average_rating, 0) as rating
		FROM
		  partners p
		  JOIN cities c ON p.city_id = c.id
		  LEFT JOIN (
			SELECT
			  partner_id,
			  AVG(rating) AS average_rating
			FROM
			  reviews
			GROUP BY
			  partner_id
		  ) r ON p.id = r.partner_id
		WHERE c.name = ? AND p.gender = ?;
		`

		rows, _ := q.db.QueryxContext(ctx, qr, daerah, jenisKelamin)
		for rows.Next() {
			var qq query.CariPasangan
			rows.StructScan(&qq)

			results = append(results, qq)
		}

		return results

	case daerah != "":
		qr := `
		SELECT
		  p.id,
		  p.name,
		  p.price,
		  c.name as c_name,
		  COALESCE(average_rating, 0) as rating
		FROM
		  partners p
		  JOIN cities c ON p.city_id = c.id
		  LEFT JOIN (
			SELECT
			  partner_id,
			  AVG(rating) AS average_rating
			FROM
			  reviews
			GROUP BY
			  partner_id
		  ) r ON p.id = r.partner_id
		WHERE c.name = ?;
		`

		rows, _ := q.db.QueryxContext(ctx, qr, daerah)
		for rows.Next() {
			var qq query.CariPasangan
			rows.StructScan(&qq)

			results = append(results, qq)
		}

		return results

	case jenisKelamin != "":
		qr := `
		SELECT
		  p.id as id,
		  p.name as name,
		  p.price as price,
		  c.name as c_name,
		  COALESCE(average_rating, 0) as rating
		FROM
		  partners p
		  JOIN cities c ON p.city_id = c.id
		  LEFT JOIN (
			SELECT
			  partner_id,
			  AVG(rating) AS average_rating
			FROM
			  reviews
			GROUP BY
			  partner_id
		  ) r ON p.id = r.partner_id
		WHERE p.gender = ?;
		`

		rows, _ := q.db.QueryxContext(ctx, qr, jenisKelamin)
		for rows.Next() {
			var qq query.CariPasangan
			rows.StructScan(&qq)

			results = append(results, qq)
		}

		return results

	default:
		qr := `
		SELECT
		  p.id,
		  p.name,
		  p.price,
		  c.name as c_name,
		  COALESCE(average_rating, 0) as rating
		FROM
		  partners p
		  JOIN cities c ON p.city_id = c.id
		  LEFT JOIN (
			SELECT
			  partner_id,
			  AVG(rating) AS average_rating
			FROM
			  reviews
			GROUP BY
			  partner_id
		  ) r ON p.id = r.partner_id
		`

		rows, _ := q.db.QueryxContext(ctx, qr)
		for rows.Next() {
			var qq query.CariPasangan
			rows.StructScan(&qq)

			results = append(results, qq)
		}

		return results
	}
}
