package query

type SeluruhTransaksi struct {
	NamaPartner    string `json:"partner_name" db:"name"`
	Kategori       string `json:"category" db:"cat_name"`
	TanggalBooking string `json:"booking_date" db:"booking_date"`
	Mulai          string `json:"start" db:"start_time"`
	Selesai        string `json:"end" db:"end_time"`
	OrderStatus    string `json:"order_status" db:"order_status"`
}
