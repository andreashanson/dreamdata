package mailjet

import (
	"github.com/andreashanson/dreamdata/pkg/config"
	"github.com/andreashanson/dreamdata/pkg/mail"
	"github.com/mailjet/mailjet-apiv3-go"
)

type MailRepository struct {
	Client *mailjet.Client
}

func NewMailRepo(c *config.MailjetConfig) *MailRepository {
	client := mailjet.NewMailjetClient(c.Key1, c.Key2)
	return &MailRepository{
		Client: client,
	}
}

func (mr *MailRepository) Send(e mail.Email) (mail.Email, error) {
	msg := &mailjet.InfoSendMail{
		FromEmail: "andreas.olof.hansson@gmail.com",
		FromName:  e.FromName,
		To:        e.To,
		Subject:   e.Subject,
		TextPart:  e.Content,
	}

	_, err := mr.Client.SendMail(msg)

	return e, err
}
