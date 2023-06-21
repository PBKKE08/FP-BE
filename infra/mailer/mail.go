package mailer

import (
	"context"
	"fmt"
	"github.com/jordan-wright/email"
	"github.com/rs/zerolog/log"
	"net/smtp"
)

var (
	identity       = ""
	emailServerURI = "localhost:1025"
	username       = "admin@example.com"
	password       = ""
	host           = "localhost"
)

func SetIdentity(s string) {
	identity = ""
	log.Info().Msgf("Set identity to %s", s)
}

func SetEmailServerURI(s string) {
	emailServerURI = s
	log.Info().Msgf("Set email server URI to %s", s)
}

func SetUsername(s string) {
	username = s
	log.Info().Msgf("Set username to %s", s)
}

func SetPassword(s string) {
	password = s
	log.Info().Msgf("Set password to %s", s)
}

func SetHost(s string) {
	host = s
	log.Info().Msgf("Set host to %s", s)
}

type Mailer func(ctx context.Context, from, to, subject, body string) error

func (m Mailer) Mail(ctx context.Context, from, to, subject, body string) error {
	return m(ctx, from, to, subject, body)
}

func SendEmail(ctx context.Context, from, to, subject, body string) error {
	e := &email.Email{
		To:      []string{to},
		From:    from,
		Subject: subject,
		Text:    []byte(body),
	}

	if err := e.Send(emailServerURI, smtp.PlainAuth(identity, username, password, host)); err != nil {
		return fmt.Errorf("mail: %w", err)
	}

	return nil
}
