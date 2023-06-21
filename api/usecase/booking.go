package usecase

import (
	"context"
	"time"
)

type BookingUsecase struct {
}

func (b *BookingUsecase) CreateBooking(ctx context.Context, userID string, partnerID string, bookingDay time.Time, timeStart string, timeEnd string) error {
	return nil
}
