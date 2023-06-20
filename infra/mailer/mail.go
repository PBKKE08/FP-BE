package mailer

import (
	"context"
	"fmt"
	"github.com/jordan-wright/email"
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
}

func SetEmailServerURI(s string) {
	emailServerURI = s
}

func SetUsername(s string) {
	username = s
}

func SetPassword(s string) {
	password = s
}

func SetHost(s string) {
	host = s
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
