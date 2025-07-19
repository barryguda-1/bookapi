package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
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
func pingHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	response := map[string]string{"status": "ok"}
	json.NewEncoder(w).Encode(response)
}

func main() {
	http.HandleFunc("/hello", helloHandler)
	http.HandleFunc("/ping", pingHandler)
	fmt.Println("Server is running on http://localhost:8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Server error:", err)
	}
}
