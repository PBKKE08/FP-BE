package authentication

import (
	"bytes"
	"context"
	"errors"
	firebase "firebase.google.com/go/v4"
	"firebase.google.com/go/v4/auth"
	"fmt"
	"github.com/rs/zerolog/log"
	"io"
	"net/http"
	"strings"
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

func (f *FirebaseAuth) Login(ctx context.Context, email, password string) error {
	user, err := f.client.GetUserByEmail(ctx, email)
	if err != nil {
		if !auth.IsUserNotFound(err) {
			log.Err(err)
			return errors.New("internal server error")
		}

		return errors.New("user not found")
	}

	if !user.EmailVerified {
		return errors.New("email not verified")
	}

	uri := "https://identitytoolkit.googleapis.com/v1/accounts:signInWithPassword?key=AIzaSyArXMliPXaBY8vuJl22GI9gaK7aPdGUMPM"
	payload := []byte(fmt.Sprintf(`{"email": "%s", "password": "%s"}`, email, password))

	req, err := http.NewRequestWithContext(ctx, "POST", uri, bytes.NewBuffer(payload))
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/json")

	client := http.DefaultClient
	resp, err := client.Do(req)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	if strings.Contains(string(body), "INVALID_PASSWORD") {
		return errors.New("bad credentials")
	}

	return nil
}

func (f *FirebaseAuth) Exists(ctx context.Context, email string) bool {
	_, err := f.client.GetUserByEmail(ctx, email)
	if err != nil && auth.IsUserNotFound(err) {
		return false
	}

	return true
}
