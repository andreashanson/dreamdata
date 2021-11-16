package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	primary "github.com/andreashanson/dreamdata/pkg/primary_mail"
)

func SendMailHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	var e primary.Email
	err := json.NewDecoder(r.Body).Decode(&e)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(500)
	}

	fmt.Println(e)

	em := primary.NewService()
	res, err := em.Send(e)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(500)
	}
	fmt.Println(res)
	json.NewEncoder(w).Encode(&e)

	//w.WriteHeader(200)
}

func ServeReactApp(w http.ResponseWriter, r *http.Request) {
	fs := http.FileServer(http.Dir("./frontend/build"))
	if _, err := os.Stat("./frontend/build" + r.RequestURI); os.IsNotExist(err) {
		http.StripPrefix(r.RequestURI, fs).ServeHTTP(w, r)
	} else {
		fs.ServeHTTP(w, r)
	}
}

func TestHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Testing handlers")
}
