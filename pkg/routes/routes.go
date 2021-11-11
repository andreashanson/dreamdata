package routes

import (
	"net/http"

	"github.com/andreashanson/dreamdata/pkg/handlers"
	"github.com/go-chi/chi"
)

func GetRoutes() http.Handler {
	r := chi.NewRouter()
	r.Post("/mail/send", http.HandlerFunc(handlers.SendMailHandler))
	r.Get("/mail", http.HandlerFunc(handlers.TestHandler))

	return r
}
