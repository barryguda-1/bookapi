package handlers

import (
	"fmt"
	"net/http"
)

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	//http://localhost:8080/hello?name=somename
	name := r.URL.Query().Get("name")
	if name == "" {
		name = "Guest"
	}

	message := fmt.Sprintf("Hello, %s!", name)
	if _, err := fmt.Fprintf(w, message); err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}
