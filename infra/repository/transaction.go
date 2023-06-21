package repository

import (
	"context"
	"database/sql"
	"github.com/PBKKE08/FP-BE/core/model/booking/transaction"
	"github.com/jmoiron/sqlx"
)

type TransactionRepository struct {
	db *sqlx.DB
}

func NewTransactionRepository(db *sqlx.DB) *TransactionRepository {
	return &TransactionRepository{db: db}
}

func (r *TransactionRepository) Save(ctx context.Context, tx transaction.Transaction) error {
	qr := `INSERT INTO transactions(id, order_id, price, payment_type, paid_at) values (?, ?, ?, ?, ?)`

	_, err := r.db.ExecContext(ctx, qr, tx.ID.String(), tx.Order.ID.String(), tx.Price, tx.PaymentType, sql.NullTime{})

	return err
}

func (r *TransactionRepository) SetPaid(ctx context.Context, id transaction.ID) error {
	qr := `
	UPDATE transactions
	SET paid_at = CURRENT_TIME
	WHERE id = ?;
	`

	_, err := r.db.ExecContext(ctx, qr, id.String())

	return err
}

func (r *TransactionRepository) WithDbTx(ctx context.Context) (*sqlx.Tx, error) {
	tx, err := r.db.BeginTxx(ctx, nil)
	return tx, err
}
