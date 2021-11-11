package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	primary "github.com/andreashanson/dreamdata/pkg/primary_mail"
)

func SendMailHandler(w http.ResponseWriter, r *http.Request) {
	var e primary.Email
	err := json.NewDecoder(r.Body).Decode(&e)
	if err != nil {
		fmt.Println(err)
	}

	em := primary.NewService()
	fmt.Println(e)
	em.Send(e)
}

func TestHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Testing handlers")
}
