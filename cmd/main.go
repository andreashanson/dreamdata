package main

import (
	"fmt"
	"net/http"

	"github.com/andreashanson/dreamdata/pkg/routes"
)

func main() {

	srv := &http.Server{
		Addr:    ":8000",
		Handler: routes.GetRoutes(),
	}

	err := srv.ListenAndServe()
	if err != nil {
		fmt.Println(err)
	}
}
