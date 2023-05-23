package auth

type InginLoginDTO struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type InginLogoutDTO struct {
	Token string `json:"-"`
}

type InginRegisterDTO struct {
	Name        string `json:"name"`
	Username    string `json:"username"`
	PhoneNumber string `json:"phone_number"`
	Email       string `json:"email"`
	Gender      string `json:"gender"`
}

type InginVerifikasiEmailDTO struct {
	Email string `json:"-"`
}
