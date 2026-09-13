package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
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

	// http server struct
	server := http.Server{
		Addr:         ":8000",
		Handler:      mux,
		ReadTimeout:  time.Second * 10,
		WriteTimeout: time.Second * 10,
		IdleTimeout:  time.Second * 60,
	}
	err := server.ListenAndServe() // using mux instead of nill (default ServeMux), it only allow our defined route
	if err != nil {
		log.Fatalf("Error occured while running server: %s", err)
	}
	fmt.Println("Server is running on port 8000")
}
