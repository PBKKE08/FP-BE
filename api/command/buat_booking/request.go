package buat_booking

import "time"

type Request struct {
	PenggunaID  string    `json:"-"`
	PartnerID   string    `json:"partner_id"`
	BookingDate time.Time `json:"booking_date"`
	TimeStart   string    `json:"time_start"`
	TimeEnd     string    `json:"time_end"`
	PaymentType string    `json:"payment_type"`
	Msg         string    `json:"message"`
}
