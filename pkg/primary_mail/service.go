package primary

import (
	"fmt"
	"log"

	mailjet "github.com/mailjet/mailjet-apiv3-go"
)

type Service struct {
	Client *mailjet.Client
}

func NewService() *Service {
	client := mailjet.NewMailjetClient("e7607096ff679c66e4b8fd1a8bac9766", "95b213f8754a237524cb8a747ffc7cab")

	return &Service{
		Client: client,
	}
}

func (s *Service) Send(e Email) error {
	messagesInfo := []mailjet.InfoMessagesV31{
		mailjet.InfoMessagesV31{
			From: &mailjet.RecipientV31{
				Email: e.From,
				Name:  e.FromName,
			},
			To: &mailjet.RecipientsV31{
				mailjet.RecipientV31{
					Email: e.To,
					Name:  "Andreas",
				},
			},
			Subject:  e.Subject,
			TextPart: e.Content,
			CustomID: "AppGettingStartedTest",
		},
	}

	messages := mailjet.MessagesV31{Info: messagesInfo}
	res, err := s.Client.SendMailV31(&messages)
	if err != nil {
		fmt.Println(err)
		log.Fatal(err)
	}
	fmt.Printf("Data: %+v\n", res)

	return nil
}
