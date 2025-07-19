package handlers

import (
	"net/http"
)

func SecureHandler(w http.ResponseWriter, r *http.Request) {
	//enforce only GET requests
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	w.Write([]byte("GET request accepted"))
}
