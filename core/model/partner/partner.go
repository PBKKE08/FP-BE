package partner

import (
	"github.com/PBKKE08/FP-BE/core/model/kategori"
	"github.com/PBKKE08/FP-BE/core/model/kota"
)

type Partner struct {
	ID           ID
	Nama         string
	Email        string
	NomorTelepon string
	JenisKelamin string
	Kota         kota.Kota
	Harga        string
	Kategori     kategori.Kategori
}
