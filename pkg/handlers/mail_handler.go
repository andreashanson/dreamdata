package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	primary "github.com/andreashanson/dreamdata/pkg/mailjet"
)

func SendMailHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:80")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	var e primary.Email
	err := json.NewDecoder(r.Body).Decode(&e)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(500)
	}

	em := primary.NewService()
	_, err = em.Send(e)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(500)
	}
	json.NewEncoder(w).Encode(&e)

}

func ServeReactApp(w http.ResponseWriter, r *http.Request) {
	fs := http.FileServer(http.Dir("./frontend/build"))
	if _, err := os.Stat("./frontend/build" + r.RequestURI); os.IsNotExist(err) {
		http.StripPrefix(r.RequestURI, fs).ServeHTTP(w, r)
	} else {
		fs.ServeHTTP(w, r)
	}
}
