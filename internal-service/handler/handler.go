package handler

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/robert076/stunning-octo-spork/shared"
)

func InternalHandler(w http.ResponseWriter, r *http.Request) {
	var e shared.ExpectedBody
	if err := json.NewDecoder(r.Body).Decode(&e); err != nil {
		http.Error(w, "Invalid JSON received. Required to have 'message' field.", http.StatusBadRequest)
		log.Println("Invalid JSON. Required to have 'message' field.")
		return
	}

	if err := json.NewEncoder(w).Encode(e); err != nil {
		http.Error(w, "Error writing the JSON.", http.StatusBadRequest)
		log.Println("Error writing the JSON.")
	}
}
