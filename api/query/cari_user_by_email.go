package query

type CariUserByEmail struct {
	Email        string `json:"email" db:"email"`
	Nama         string `json:"name" db:"name"`
	ID           string `json:"id" db:"id"`
	NomorTelepon string `json:"telephone" db:"telephone"`
	JenisKelamin string `json:"gender" db:"gender"`
}
