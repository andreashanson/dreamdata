package mailjet

import (
	"fmt"

	"github.com/andreashanson/dreamdata/pkg/mail"
	"github.com/mailjet/mailjet-apiv3-go"
)

type MailRepository struct {
	Client *mailjet.Client
}

func NewMailRepo() *MailRepository {
	client := mailjet.NewMailjetClient("e7607096ff679c66e4b8fd1a8bac9766", "95b213f8754a237524cb8a747ffc7cab")
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

	if err != nil {
		fmt.Println(err)
	}
	return e, nil
}
