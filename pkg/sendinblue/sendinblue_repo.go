package sendinblue

import (
	"encoding/base64"
	"fmt"
	"net/smtp"

	"github.com/andreashanson/dreamdata/pkg/config"
	"github.com/andreashanson/dreamdata/pkg/mail"
)

type SendingblueRepository struct {
	host     string
	user     string
	password string
}

func NewSendinBlueRepo(c *config.SMTPConfig) *SendingblueRepository {
	host := c.SMTPHost
	user := c.SMTPUser
	pw := c.SMTPPassword
	host = "smtp-relay.sendinblue.com"
	user = "andreas.olof.hansson@gmail.com"
	pw = "shYZSJ5mp1cLg4n2"
	return &SendingblueRepository{
		host:     host,
		user:     user,
		password: pw,
	}
}

func (sbr SendingblueRepository) Send(e mail.Email) (mail.Email, error) {

	auth := smtp.PlainAuth(
		"",
		sbr.user,
		sbr.password,
		sbr.host,
	)

	header := make(map[string]string)
	header["From"] = sbr.user
	header["To"] = e.To
	header["Subject"] = e.Subject
	header["MIME-Version"] = "1.0"
	header["Content-Type"] = "text/plain; charset=\"utf-8\""
	header["Content-Transfer-Encoding"] = "base64"

	body := e.Content

	message := ""
	for k, v := range header {
		message += fmt.Sprintf("%s: %s\r\n", k, v)
	}
	message += "\r\n" + base64.StdEncoding.EncodeToString([]byte(body))

	err := smtp.SendMail(
		sbr.host+":587",
		auth,
		sbr.user,
		[]string{e.To},
		[]byte(message),
	)

	return e, err
}
