package terima_pembayaran

import (
	"context"
	"github.com/PBKKE08/FP-BE/core/model/booking/transaction"
	"github.com/PBKKE08/FP-BE/core/repository"
)

type TerimaPembayaran struct {
	TransactionRepo repository.Transaction
}

func (t *TerimaPembayaran) Execute(ctx context.Context, r Request) error {
	txID, err := transaction.NewIDFrom(r.TxID)
	if err != nil {
		return err
	}

	err = t.TransactionRepo.SetPaid(ctx, txID)
	if err != nil {
		return err
	}

	return nil
}
