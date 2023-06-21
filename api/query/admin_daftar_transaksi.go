package query

type TransaksiNonTerbayar struct {
	NamaPembooking string `json:"booker_name" db:"booker_name"`
	PartnerID      string `json:"partner_id" db:"partner_id"`
	Email          string `json:"partner_email" db:"partner_email"`
	Nama           string `json:"partner_name" db:"partner_name"`
	WaktuBooking   string `json:"booking_date" db:"booking_date"`
	Mulai          string `json:"time_start" db:"time_start"`
	Selesai        string `json:"time_end" db:"time_end"`
	TxID           string `json:"transaction_id" db:"tx_id"`
	Harga          string `json:"price" db:"tx_price"`
	PaymentTipe    string `json:"payment_type" db:"payment_type"`
}
