package buat_partner

type Request struct {
	ID           string `json:"id"`
	Nama         string `json:"name"`
	Email        string `json:"email"`
	NomorTelepon string `json:"telephone"`
	JenisKelamin string `json:"gender"`
	KategoriID   string `json:"category_id"`
	Price        string `json:"price"`
	KotaID       string `json:"city_id"`
}
