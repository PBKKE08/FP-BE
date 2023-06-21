package pengguna

import "github.com/PBKKE08/FP-BE/core/model/kota"

type Pengguna struct {
	ID           ID
	Nama         string
	Email        string
	NomorTelepon string
	JenisKelamin string
	Kota         kota.Kota
}
