package terima_pembayaran

type Request struct {
	NamaPembooking string `json:"booker_name"`
	PartnerID      string `json:"partner_id"`
	Email          string `json:"partner_email"`
	Nama           string `json:"partner_name"`
	WaktuBooking   string `json:"booking_date"`
	Mulai          string `json:"time_start"`
	Selesai        string `json:"time_end"`
	TxID           string `json:"transaction_id"`
	Harga          string `json:"price"`
	PaymentTipe    string `json:"payment_type"`
}
