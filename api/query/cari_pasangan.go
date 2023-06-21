package query

type CariPasangan struct {
	ID           string  `json:"id" db:"id"`
	Nama         string  `json:"nama" db:"name"`
	Harga        string  `json:"harga" db:"price"`
	Rating       float64 `json:"rating" db:"rating"`
	Daerah       string  `json:"daerah" db:"c_name"`
	Kategori     string  `json:"kategori" db:"cat_name"`
	JenisKelamin string  `json:"gender" db:"gender"`
}
