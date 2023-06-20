package authentication

import (
	"context"
	"errors"
	firebase "firebase.google.com/go/v4"
	"firebase.google.com/go/v4/auth"
	"github.com/rs/zerolog/log"
)

type FirebaseAuth struct {
	client *auth.Client
}

func NewFirebaseAuth(app *firebase.App) (*FirebaseAuth, error) {
	client, err := app.Auth(context.Background())
	if err != nil {
		return nil, err
	}

	return &FirebaseAuth{client: client}, nil
}

func (f *FirebaseAuth) Register(ctx context.Context, email, password string) (string, error) {
	user, err := f.client.GetUserByEmail(ctx, email)
	if err != nil {
		if !auth.IsUserNotFound(err) {
			log.Err(err)
			return "", errors.New("internal server error")
		}

		params := (&auth.UserToCreate{}).Email(email).Password(password)
		_, err := f.client.CreateUser(ctx, params)
		if err != nil {
			log.Err(err)
			return "", errors.New("internal server error")
		}

		link, err := f.client.EmailVerificationLink(ctx, email)

		return link, err
	}

	if user.EmailVerified {
		return "", errors.New("user already exists")
	}

	link, err := f.client.EmailVerificationLink(ctx, email)

	return link, err
}
