package transaction

import (
	"github.com/PBKKE08/FP-BE/core/model/booking/order"
	"time"
)

type Transaction struct {
	ID          ID
	Order       order.Order
	Price       string
	PaymentType string
	PaidAt      time.Time
}

func (t Transaction) IsPaid() bool {
	zeroTime := time.Time{}

	return t.PaidAt != zeroTime
}
