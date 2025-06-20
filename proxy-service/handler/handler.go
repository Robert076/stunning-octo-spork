package handler

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"time"
)

type expectedBody struct {
	Message string `json:"message"`
}

func ProxyHandler(w http.ResponseWriter, r *http.Request) {
	var e expectedBody
	if err := json.NewDecoder(r.Body).Decode(&e); err != nil {
		http.Error(w, "Invalid JSON. Required to have 'message' field.", http.StatusBadRequest)
		log.Println("Invalid JSON. Required to have 'message' field.")
		return
	}

	marshalled, err := json.Marshal(e)
	if err != nil {
		http.Error(w, "Error encoding JSON.", http.StatusBadRequest)
		log.Println("Error encoding JSON.")
		return
	}

	internalServiceHost := os.Getenv("INTERNAL_SERVICE_HOST")
	if internalServiceHost == "" {
		http.Error(w, "Setup issue. Check backend logs.", http.StatusBadRequest)
		log.Println("Setup issue. No env var for INTERNAL_SERVICE_HOST")
		return
	}

	req, err := http.NewRequest("POST", internalServiceHost, bytes.NewReader(marshalled))
	if err != nil {
		http.Error(w, "Impossible to build request", http.StatusBadRequest)
		log.Println("Impossible to build request")
		return
	}

	req.Header.Set("Content-Type", "application/json")

	client := http.Client{Timeout: 10 * time.Second}

	_, err = client.Do(req)
	if err != nil {
		http.Error(w, "Cannot reach internal service.", http.StatusBadRequest)
		log.Println("Cannot reach internal service")
		return
	}
}
