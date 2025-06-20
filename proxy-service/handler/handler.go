package handler

import (
	"log"
	"net/http"
)

func ProxyHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Hello, World!")
}
