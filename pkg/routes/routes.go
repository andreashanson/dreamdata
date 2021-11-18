package routes

import (
	"net/http"

	"github.com/andreashanson/dreamdata/pkg/handlers"
	"github.com/go-chi/chi"
)

func GetRoutes() http.Handler {

	r := chi.NewRouter()
	r.Post("/mail/send", http.HandlerFunc(handlers.SendMailHandler))
	r.Get("/*", http.HandlerFunc(handlers.ServeReactApp))
	return r
}
