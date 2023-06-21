package query

type DetailPartner struct {
	ID           string  `json:"id" db:"id"`
	Nama         string  `json:"name" db:"name"`
	Kota         string  `json:"city" db:"city_name"`
	Harga        string  `json:"price" db:"price"`
	Email        string  `json:"email" db:"email"`
	Kategori     string  `json:"category" db:"cat_name"`
	Rating       float64 `json:"rating" db:"rating"`
	Deskripsi    string  `json:"description" db:"description"`
	JenisKelamin string  `json:"gender" db:"gender"`
}
