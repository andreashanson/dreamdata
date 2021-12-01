package main

import (
	"fmt"
	"net/http"

	"github.com/andreashanson/dreamdata/pkg/config"
	"github.com/andreashanson/dreamdata/pkg/handlers"
	"github.com/andreashanson/dreamdata/pkg/mail"
	"github.com/andreashanson/dreamdata/pkg/mailjet"
	"github.com/andreashanson/dreamdata/pkg/routes"
	"github.com/andreashanson/dreamdata/pkg/sendinblue"
)

func main() {

	cfg := config.NewConfig()
	mailjetConfig := cfg.MailjetConfig
	sendinblueConfig := cfg.SMTPConfig

	mailjetRepo := mailjet.NewMailjetRepo(mailjetConfig)
	mailSrv1 := mail.NewService(mailjetRepo)

	sendinblueRepo := sendinblue.NewSendinBlueRepo(sendinblueConfig)
	mailSrv2 := mail.NewService(sendinblueRepo)

	handlerRepo := handlers.NewHandlersRepo(mailSrv1, mailSrv2)

	newRoutes := routes.GetRoutes(handlerRepo)

	srv := &http.Server{
		Addr:    ":8000",
		Handler: newRoutes,
	}

	err := srv.ListenAndServe()
	if err != nil {
		fmt.Println(err)
	}
}
