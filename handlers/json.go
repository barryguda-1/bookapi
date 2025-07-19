package handlers

import (
	"encoding/json"
	"net/http"
)

type Message struct {
	Text string `json:"text"`
}

func JsonHandler(w http.ResponseWriter, r *http.Request) {
	msg := Message{Text: "Welcome to the Go API"}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(msg)
}
