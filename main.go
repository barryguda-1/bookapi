package main

import (
	"bookapi/handlers"
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/hello", handlers.HelloHandler)
	http.HandleFunc("/ping", handlers.PingHandler)
	http.HandleFunc("/welcome", handlers.JsonHandler)
	http.HandleFunc("/secure", handlers.SecureHandler)
	fmt.Println("Server is running on http://localhost:8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Server error:", err)
	}
}
