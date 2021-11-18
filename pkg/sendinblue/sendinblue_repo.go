package sendinblue

import (
	"log"
	"net/smtp"

	"github.com/andreashanson/dreamdata/pkg/mail"
)

type SendingblueRepository struct {
	auth *smtp.Auth
}

func NewSendinBlueRepo() *SendingblueRepository {
	pw := "shYZSJ5mp1cLg4n2"

	auth := smtp.PlainAuth("", "andreas.olof.hansson@gmail.com", pw, "smtp-relay.sendinblue.com")

	return &SendingblueRepository{
		auth: &auth,
	}
}

func (sbr SendingblueRepository) Send(e mail.Email) (mail.Email, error) {

	// Set up authentication information.

	// Connect to the server, authenticate, set the sender and recipient,
	// and send the email all in one step.
	to := []string{e.To}
	msg := []byte("To: recipient@example.net\r\n" +
		"Subject: discount Gophers!\r\n" +
		"\r\n" +
		"This is the email body.\r\n")
	err := smtp.SendMail("smtp-relay.sendinblue.com:587", *sbr.auth, e.From, to, msg)
	if err != nil {
		log.Fatal(err)
	}
	return e, nil
}
