package main

import (
	"encoding/json"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	response := map[string]string{"message": "hello world"}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func main() {
	http.HandleFunc("/hello", helloHandler)
	http.ListenAndServe(":8080", nil)
}
