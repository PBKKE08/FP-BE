package query

type CariPartnerByEmail struct {
	Email        string `json:"email" db:"email"`
	Nama         string `json:"name" db:"name"`
	ID           string `json:"id" db:"id"`
	NomorTelepon string `json:"telephone" db:"telephone"`
	JenisKelamin string `json:"gender" db:"gender"`
	IsApproved   bool   `json:"is_approved" db:"is_approved"`
}
