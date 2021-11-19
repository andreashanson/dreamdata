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
	key1 := c.Key1
	key2 := c.Key2
	key1 = "e7607096ff679c66e4b8fd1a8bac9766"
	key2 = "95b213f8754a237524cb8a747ffc7cab"
	client := mailjet.NewMailjetClient(key1, key2)
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
