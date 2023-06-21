package usecase

import (
	"context"
	"errors"
	"github.com/PBKKE08/FP-BE/api/command/buat_booking"
	"github.com/PBKKE08/FP-BE/core/model/booking"
	"time"
)

type BuatBookingCmd interface {
	Execute(ctx context.Context, req buat_booking.Request, userID string) error
}

type BookingUsecase struct {
	buatbookingCmd BuatBookingCmd
}

func NewBookingUsecase(buatbookingCmd BuatBookingCmd) *BookingUsecase {
	return &BookingUsecase{buatbookingCmd: buatbookingCmd}
}

func (b *BookingUsecase) CreateBooking(ctx context.Context, userID string, req buat_booking.Request) error {
	if !isValidTime(req.TimeStart) {
		return errors.New("time start is not valid")
	}

	if !isValidTime(req.TimeEnd) {
		return errors.New("time end is not valid")
	}

	if !booking.IsPaymentTypeValid(req.PaymentType) {
		return errors.New("payment type is not valid")
	}

	err := b.buatbookingCmd.Execute(ctx, req, userID)
	return err
}

func isValidTime(t string) bool {
	_, err := time.Parse("15:04", t)
	return err == nil
}
