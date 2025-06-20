package main

import (
	"log"
	"net/http"

	"github.com/robert076/stunning-octo-spork/internal-service/handler"
)

func main() {
	http.HandleFunc("/internal", handler.InternalHandler)

	if err := http.ListenAndServe(":5433", nil); err != nil {
		log.Println("Error starting http server.")
	}
}
