package buat_booking

import (
	"context"
	"fmt"
	"github.com/PBKKE08/FP-BE/core/model/booking/order"
	"github.com/PBKKE08/FP-BE/core/model/booking/transaction"
	"github.com/PBKKE08/FP-BE/core/model/partner"
	"github.com/PBKKE08/FP-BE/core/model/pengguna"
	"github.com/PBKKE08/FP-BE/core/repository"
	"strconv"
	"time"
)

type BuatBooking struct {
	TransactionRepo repository.Transaction
	OrderRepo       repository.Order
	PenggunaRepo    repository.Pengguna
	PartnerRepo     repository.Partner
}

func (b *BuatBooking) Execute(ctx context.Context, req Request, userID string) error {
	partnerID, err := partner.NewIDFrom(req.PartnerID)
	if err != nil {
		return err
	}

	penggunaID, err := pengguna.NewIDFrom(userID)
	if err != nil {
		return err
	}

	partnerData, err := b.PartnerRepo.ByID(ctx, partnerID)
	if err != nil {
		return err
	}

	penggunaData, err := b.PenggunaRepo.ByID(ctx, penggunaID)
	if err != nil {
		return err
	}

	orderData := order.Order{
		ID:         order.NewID(),
		BookingDay: req.BookingDate,
		TimeStart:  req.TimeStart,
		TimeEnd:    req.TimeEnd,
		Message:    req.Msg,
		Pengguna:   penggunaData,
		Partner:    partnerData,
	}

	err = b.OrderRepo.Save(ctx, orderData)
	if err != nil {
		return fmt.Errorf("error at saving order: %w", err)
	}

	parsedHarga, err := strconv.ParseFloat(partnerData.Harga, 64)
	if err != nil {
		return err
	}

	price := strconv.FormatFloat(parsedHarga*orderData.GetDuration(), 'f', -1, 64)

	tx := transaction.Transaction{
		ID:          transaction.NewID(),
		Order:       orderData,
		Price:       price,
		PaymentType: req.PaymentType,
		PaidAt:      time.Time{},
	}

	err = b.TransactionRepo.Save(ctx, tx)
	if err != nil {
		return fmt.Errorf("error at saving tx: %w", err)
	}

	return err
}
