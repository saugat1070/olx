package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {

	mux := http.NewServeMux()
	mux.HandleFunc("GET /health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK) // go provide http status code as constant
		w.Header().Set("content-type", "application/json")
		w.Write([]byte(`{"status": "okay"}`)) // conversion from string to byte
	})

	// http.HandleFunc("GET /health", func(w http.ResponseWriter, r *http.Request) {
	// 	w.WriteHeader(http.StatusOK) // go provide http status code as constant
	// 	w.Header().Set("content-type", "application/json")
	// 	w.Write([]byte(`{"status": "okay"}`)) // conversion from string to byte
	// })
	// @description: adding default http DefaultServeMux can cause script injection on our server
	// so we use mux from http package instead of default ServeMux

	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		log.Fatalf("Error occured while running server: %s", err)
	}
	fmt.Println("Server is running on port 8000")
}
