package usecase

import (
	"context"
	"errors"
	"github.com/PBKKE08/FP-BE/api/command/buat_user"
	"github.com/PBKKE08/FP-BE/api/query"
)

type AuthProvider interface {
	Register(ctx context.Context, email, password string) (string, error)
	Login(ctx context.Context, email, password string) error
}

type Mailer interface {
	Mail(ctx context.Context, from, to, subject, body string) error
}

type BuatUserCommand interface {
	Execute(ctx context.Context, request buat_user.BuatUserRequest) error
}

type CariUserQuery interface {
	ByUserEmail(ctx context.Context, email string) query.CariUserByEmail
}

type JWTProvider interface {
	GenerateToken(email, nama, id, nomorTelepon, jenisKelamin string) string
}

type AuthUsecase struct {
	buatUser     BuatUserCommand
	authProvider AuthProvider
	mailer       Mailer
	cariUser     CariUserQuery
	jwtProvider  JWTProvider
}

func NewAuthUsecase(buatUser BuatUserCommand, authProvider AuthProvider, cariUser CariUserQuery, mailer Mailer, jwtProvider JWTProvider) *AuthUsecase {
	return &AuthUsecase{buatUser: buatUser, authProvider: authProvider, mailer: mailer, cariUser: cariUser, jwtProvider: jwtProvider}
}

func (a *AuthUsecase) Register(ctx context.Context, req buat_user.BuatUserRequest) error {
	if req.Email == "" {
		return errors.New("empty email")
	}

	if req.Nama == "" {
		return errors.New("empty name")
	}

	if req.Telepon == "" {
		return errors.New("empty phone number")
	}

	if req.KotaID == "" {
		return errors.New("empty city")
	}

	if req.JenisKelamin != "" && !(req.JenisKelamin == "f" || req.JenisKelamin == "m") {
		return errors.New("unknown gender")
	}

	err := a.buatUser.Execute(ctx, req)
	if err != nil {
		return err
	}

	link, err := a.authProvider.Register(ctx, req.Email, req.Password)
	if err != nil {
		return err
	}

	err = a.mailer.Mail(ctx, "testing@gmail.com", req.Email, "Register", link)
	if err != nil {
		return err
	}

	return nil
}

func (a *AuthUsecase) Login(ctx context.Context, email, password string) (string, error) {
	err := a.authProvider.Login(ctx, email, password)
	if err != nil {
		return "", err
	}

	result := a.cariUser.ByUserEmail(ctx, email)

	token := a.jwtProvider.GenerateToken(result.Email, result.Nama, result.ID, result.NomorTelepon, result.JenisKelamin)

	return token, nil
}
