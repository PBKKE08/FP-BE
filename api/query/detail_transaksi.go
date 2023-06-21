package query

type DetailTransaksi struct {
	OrderID        string `json:"order_id" db:"order_id"`
	NamaPartner    string `json:"partner_name" db:"name"`
	PartnerID      string `json:"partner_id" db:"partner_id"`
	Kategori       string `json:"category" db:"cat_name"`
	TanggalBooking string `json:"booking_date" db:"booking_date"`
	Mulai          string `json:"start" db:"start_time"`
	Selesai        string `json:"end" db:"end_time"`
	OrderStatus    string `json:"order_status" db:"order_status"`
	Harga          string `json:"price" db:"price"`
	PaymentType    string `json:"payment_type" db:"payment_type"`
	SudahDibayar   bool   `json:"is_paid" db:"is_paid"`
}
