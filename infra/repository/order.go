package repository

import (
	"context"
	"github.com/PBKKE08/FP-BE/core/model/booking/order"
	"github.com/jmoiron/sqlx"
)

type OrderRepo struct {
	db *sqlx.DB
}

func NewOrderRepo(db *sqlx.DB) *OrderRepo {
	return &OrderRepo{db: db}
}

func (r *OrderRepo) Save(ctx context.Context, order order.Order) error {
	qr := `
	INSERT INTO orders(id, booking_day, time_start, time_end, message, partners_id, user_id) values (?, ?, ?, ?, ?, ?, ?);
	`

	_, err := r.db.ExecContext(ctx,
		qr,
		order.ID.String(),
		order.ToBookingTypeString(),
		order.TimeStart,
		order.TimeEnd,
		order.Message,
		order.Partner.ID.String(),
		order.Pengguna.ID.String(),
	)

	return err
}

func (r *OrderRepo) ByID(ctx context.Context, id order.ID) (order.Order, error) {
	return order.Order{}, nil
}
