package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/robert076/stunning-octo-spork/proxy-service/handler"
)

func main() {
	internalServiceHost := os.Getenv("INTERNAL_SERVICE_HOST")
	if internalServiceHost == "" {
		log.Println("Setup issue. No env var for INTERNAL_SERVICE_HOST.")
		return
	}

	http.HandleFunc("/proxy", handler.ProxyHandler)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("Error setting up http server")
	}
}
