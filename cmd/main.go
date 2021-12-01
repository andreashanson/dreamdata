package main

import (
	"fmt"

	"github.com/andreashanson/dreamdata/pkg/config"
	"github.com/andreashanson/dreamdata/pkg/handlers"
	"github.com/andreashanson/dreamdata/pkg/mail"
	"github.com/andreashanson/dreamdata/pkg/mailjet"
	"github.com/andreashanson/dreamdata/pkg/routes"
	"github.com/andreashanson/dreamdata/pkg/sendinblue"
	"github.com/andreashanson/dreamdata/pkg/server"
)

func main() {

	cfg := config.NewConfig()

	fmt.Println(cfg.MailjetConfig)
	mailjetConfig := cfg.MailjetConfig
	sendinblueConfig := cfg.SMTPConfig

	mailjetRepo := mailjet.NewMailjetRepo(mailjetConfig)
	mailSrv1 := mail.NewService(mailjetRepo)

	sendinblueRepo := sendinblue.NewSendinBlueRepo(sendinblueConfig)
	mailSrv2 := mail.NewService(sendinblueRepo)

	handlerRepo := handlers.NewHandlersRepo(mailSrv1, mailSrv2)

	r := routes.GetRoutes(handlerRepo)

	srv := server.NewServer(":8000", r)

	err := srv.Run()
	if err != nil {
		fmt.Println(err)
	}
}
