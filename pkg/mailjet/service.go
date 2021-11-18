package mailjet

import (
	"fmt"

	mailjet "github.com/mailjet/mailjet-apiv3-go"
)

type Service struct {
	Client  *mailjet.Client
	Client2 *mailjet.Client
}

func NewService() *Service {
	client := mailjet.NewMailjetClient("e7607096ff679c66e4b8fd1a8bac9766", "95b213f8754a237524cb8a747ffc7cab")
	client2 := mailjet.NewMailjetClient("e7607096ff679c66e4b8fd1a8bac9766", "95b213f8754a237524cb8a747ffc7cab")

	return &Service{
		Client:  client,
		Client2: client2,
	}
}

func (s *Service) Send(e Email) (*mailjet.SentResult, error) {

	msg := &mailjet.InfoSendMail{
		FromEmail: "andreas.olof.hansson@gmail.com",
		FromName:  e.FromName,
		To:        e.To,
		Subject:   e.Subject,
		TextPart:  e.Content,
	}

	res, err := s.Client.SendMail(msg)
	if err != nil {
		fmt.Println("COULD NOT SEND EMAIL")
		fmt.Println("TRYING WITH CLIENT 2")
		res, err = s.Client2.SendMail(msg)
	}
	fmt.Printf("Data: %+v\n", res)

	return res, err
}
