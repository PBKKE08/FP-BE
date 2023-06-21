package repository

import (
	"context"
	"errors"

	"github.com/PBKKE08/FP-BE/core/model/review"
	"github.com/jmoiron/sqlx"
)

type ReviewRepository struct {
	db *sqlx.DB
}

func NewReviewRepository(db *sqlx.DB) *ReviewRepository {
	return &ReviewRepository{db}
}

func (r *ReviewRepository) Save(ctx context.Context, review review.Review) error {
	q := `INSERT INTO reviews(id, user_id, partner_id, rating, comment) VALUES(?, ?, ?, ?, ?)`

	result, err := r.db.ExecContext(ctx, q, review.ID.String(), review.Pengguna.ID.String(), review.Partner.ID.String(), review.Rating, review.Comment)
	if err != nil {
		return err
	}

	if affected, _ := result.RowsAffected(); affected == 0 {
		return errors.New("internal server error")
	}

	return nil
}
