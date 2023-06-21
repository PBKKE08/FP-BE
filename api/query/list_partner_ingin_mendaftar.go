package query

type PartnerInginMendaftar struct {
	PartnerID    string `json:"partner_id" db:"partner_id"`
	Nama         string `json:"name" db:"name"`
	JenisKelamin string `json:"gender" db:"gender"`
	Harga        string `json:"price" db:"price"`
	Kategori     string `json:"category_name" db:"category_name"`
	Description  string `json:"description" db:"description"`
	Kota         string `json:"city" db:"city_name"`
}
