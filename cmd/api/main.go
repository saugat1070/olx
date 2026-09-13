package main

import (
	"log"
	"net/http"
	"time"

	"github.com/saugat1070/olx-api/internal/config"
	"github.com/saugat1070/olx-api/internal/handlers"
)

func main() {
	cfg := config.MustLoadConfig()
	mux := http.NewServeMux()
	mux.HandleFunc("GET /healthz", handlers.Health)

	// http server struct
	server := http.Server{
		Addr:         ":" + cfg.PORT,
		Handler:      mux,
		ReadTimeout:  time.Second * 10,
		WriteTimeout: time.Second * 30,
		IdleTimeout:  time.Second * 60,
	}
	log.Printf("Server is running on port :%s", cfg.PORT)
	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}

}
