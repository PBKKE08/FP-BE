package buat_user

type BuatUserRequest struct {
	Nama         string `json:"name"`
	Telepon      string `json:"phone_number"`
	Email        string `json:"email"`
	JenisKelamin string `json:"gender"`
	KotaID       string `json:"city_id"`
	Password     string `json:"password"`
}
