package buat_partner

type Request struct {
	Nama         string `json:"name"`
	Email        string `json:"email"`
	Password     string `json:"password"`
	NomorTelepon string `json:"telephone"`
	JenisKelamin string `json:"gender"`
	KategoriID   string `json:"category_id"`
	Harga        string `json:"price"`
	KotaID       string `json:"city_id"`
	Description  string `json:"description"`
}
