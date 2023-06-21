package authentication

import (
	"github.com/golang-jwt/jwt/v5"
	"time"
)

var (
	secretKey = "123123123"
)

func SetSecretKey(s string) {
	secretKey = s
}

type CustomClaims struct {
	Email        string `json:"email"`
	Nama         string `json:"name"`
	ID           string `json:"id"`
	NomorTelepon string `json:"telephone"`
	JenisKelamin string `json:"gender"`
	jwt.RegisteredClaims
}

type JWTProvider func(email, nama, id, nomorTelepon, jenisKelamin string) string

func (j JWTProvider) GenerateToken(email, nama, id, nomorTelepon, jenisKelamin string) string {
	return j(email, nama, id, nomorTelepon, jenisKelamin)
}

func GenerateToken(email, nama, id, nomorTelepon, jenisKelamin string) string {
	claims := &CustomClaims{
		Email:        email,
		Nama:         nama,
		ID:           id,
		NomorTelepon: nomorTelepon,
		JenisKelamin: jenisKelamin,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 72)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, _ := token.SignedString([]byte(secretKey))

	return t
}
