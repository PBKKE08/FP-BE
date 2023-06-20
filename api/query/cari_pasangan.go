package query

type CariPasangan struct {
	ID     string `json:"id,omitempty" db:"id"`
	Nama   string `json:"nama,omitempty" db:"name"`
	Harga  string `json:"harga,omitempty" db:"price"`
	Rating int    `json:"rating,omitempty" db:"rating"`
	Daerah string `json:"daerah,omitempty" db:"c_name"`
}
