package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/andreashanson/dreamdata/pkg/mail"
)

type HandlersRepo struct {
	MailSrv1 *mail.Service
	MailSrv2 *mail.Service
}

func NewHandlersRepo(m1 *mail.Service, m2 *mail.Service) *HandlersRepo {
	return &HandlersRepo{
		MailSrv1: m1,
		MailSrv2: m2,
	}
}

func (h *HandlersRepo) SendMailHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:80")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	var e mail.Email
	err := json.NewDecoder(r.Body).Decode(&e)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(500)
		return
	}

	_, err = h.MailSrv1.Send(e)
	if err != nil {
		fmt.Println("Error first mail service: ", err)
		fmt.Println("Could not send with first mail service.")
		fmt.Println("Try and send with second mail service.")
		_, err = h.MailSrv2.Send(e)
		if err != nil {
			fmt.Println("Could not send with second mail service")
			fmt.Println("Error second mail service: ", err)
			w.WriteHeader(500)
			return
		}
	}
	err = json.NewEncoder(w).Encode(&e)
	if err != nil {
		fmt.Println(err)
	}
}

func ServeReactApp(w http.ResponseWriter, r *http.Request) {
	fs := http.FileServer(http.Dir("./frontend/build"))
	if _, err := os.Stat("./frontend/build" + r.RequestURI); os.IsNotExist(err) {
		http.StripPrefix(r.RequestURI, fs).ServeHTTP(w, r)
	} else {
		fs.ServeHTTP(w, r)
	}
}
