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
	return &SendingblueRepository{
		host:     c.SMTPHost,
		user:     c.SMTPUser,
		password: c.SMTPPassword,
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
