package main

import (
	"encoding/json"
	"net/http"
	"time"
)

type ContactInfo struct {
	Name      string
	Email     string
	CreatedAt time.Time
}

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/contactInfo", contactInfoHandler)

	http.ListenAndServe(":8080", mux)
}

func contactInfoHandler(w http.ResponseWriter, r *http.Request) {
	contactInfo := ContactInfo{}

	err := json.NewDecoder(r.Body).Decode(&contactInfo)
	if err != nil {
		panic(err)
	}

	contactInfo.CreatedAt = time.Now().Local()

	contactInfoJSON, err := json.Marshal(contactInfo)
	if err != nil {
		panic(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(contactInfoJSON)
}
