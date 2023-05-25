package pkg

import (
	"errors"
	"github.com/golang-jwt/jwt"
	"time"
)

var KEY []byte = []byte("the secret of the sun")
var ErrCantGenerateToken = errors.New("cant generate token")

type JWTClaims struct {
	jwt.StandardClaims
	Name      string `json:"name"`
	Email     string `json:"email"`
	Gender    string `json:"gender"`
	Telephone string `json:"telephone"`
}

type JWTGenerator struct{}

func (JWTGenerator) GenerateToken(id, name, email, telephone, gender string) (string, error) {
	claims := JWTClaims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Duration(1) * time.Hour).Unix(),
			Id:        id,
			Issuer:    "socium-rentalis",
		},
		Name:      name,
		Email:     email,
		Gender:    gender,
		Telephone: telephone,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString(KEY)
	if err != nil {
		return "", ErrCantGenerateToken
	}

	return signedToken, nil
}
