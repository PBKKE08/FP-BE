package order

import (
	"time"
)

type Order struct {
	ID         ID
	BookingDay time.Time
	TimeStart  string
	TimeEnd    string
	Message    string
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
