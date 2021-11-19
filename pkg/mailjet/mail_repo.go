package mailjet

import (
	"fmt"

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
	fmt.Println(key1)
	fmt.Println(key2)
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
