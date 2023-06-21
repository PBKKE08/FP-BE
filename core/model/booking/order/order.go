package order

import (
	"github.com/PBKKE08/FP-BE/core/model/partner"
	"github.com/PBKKE08/FP-BE/core/model/pengguna"
	"time"
)

type Order struct {
	ID         ID
	BookingDay time.Time
	TimeStart  string
	TimeEnd    string
	Message    string
	Partner    partner.Partner
	Pengguna   pengguna.Pengguna
}

func (o Order) GetDuration() float64 {
	timeStart, _ := time.Parse("15:04", o.TimeStart)
	timeEnd, _ := time.Parse("15:04", o.TimeEnd)
	duration := timeEnd.Sub(timeStart)

	return duration.Minutes()
}

func (o Order) IsTimeValid() bool {
	timeStart, _ := time.Parse("15:04", o.TimeStart)
	timeEnd, _ := time.Parse("15:04", o.TimeEnd)

	if timeStart.After(timeEnd) {
		return false
	}

	return true
}

func (o Order) ToBookingTypeString() string {
	return o.BookingDay.Format("2006-01-02")
}
