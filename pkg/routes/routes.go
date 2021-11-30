package routes

import (
	"net/http"

	"github.com/andreashanson/dreamdata/pkg/config"
	"github.com/andreashanson/dreamdata/pkg/handlers"
	"github.com/andreashanson/dreamdata/pkg/mail"
	"github.com/andreashanson/dreamdata/pkg/mailjet"
	"github.com/andreashanson/dreamdata/pkg/sendinblue"
	"github.com/go-chi/chi"
)

func GetRoutes() http.Handler {
	cfg := config.NewConfig()
	mailjetConfig := cfg.MailjetConfig
	sendinblueConfig := cfg.SMTPConfig

	mailjetRepo := mailjet.NewMailjetRepo(mailjetConfig)
	mailSrv1 := mail.NewService(mailjetRepo)

	sendinblueRepo := sendinblue.NewSendinBlueRepo(sendinblueConfig)
	mailSrv2 := mail.NewService(sendinblueRepo)

	handlerRepo := handlers.NewHandlersRepo(mailSrv1, mailSrv2)

	r := chi.NewRouter()
	r.Post("/mail/send", http.HandlerFunc(handlerRepo.SendMailHandler))
	r.Get("/*", http.HandlerFunc(handlers.ServeReactApp))
	return r
}
