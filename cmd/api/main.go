package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/saugat1070/olx-api/internal/config"
)

func main() {
	cfg := config.MustLoadConfig()
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
		Addr:         ":" + cfg.PORT,
		Handler:      mux,
		ReadTimeout:  time.Second * 10,
		WriteTimeout: time.Second * 10,
		IdleTimeout:  time.Second * 60,
	}
	log.Printf("Server is running on port :%s", os.Getenv("PORT"))
	err := server.ListenAndServe() // using mux instead of nill (default ServeMux), it only allow our defined route
	if err != nil {
		panic(err)
	}

}
