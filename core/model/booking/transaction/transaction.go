package transaction

import (
	"time"
)

type Transaction struct {
	ID          ID
	Price       string
	PaymentType string
	PaidAt      time.Time
}

func (t Transaction) IsPaid() bool {
	zeroTime := time.Time{}

	return t.PaidAt != zeroTime
}
