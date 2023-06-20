package usecase

import (
	"context"
	"errors"
	"github.com/PBKKE08/FP-BE/api/command/buat_user"
)

type AuthProvider interface {
	Register(ctx context.Context, email, password string) (string, error)
}

type Mailer interface {
	Mail(ctx context.Context, from, to, subject, body string) error
}

type BuatUserCommand interface {
	Execute(ctx context.Context, request buat_user.BuatUserRequest) error
}

type AuthUsecase struct {
	buatUser     BuatUserCommand
	authProvider AuthProvider
	mailer       Mailer
}

func NewAuthUsecase(buatUser BuatUserCommand, authProvider AuthProvider, mailer Mailer) *AuthUsecase {
	return &AuthUsecase{buatUser: buatUser, authProvider: authProvider, mailer: mailer}
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
