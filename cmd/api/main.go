package main

import (
	"log"
	"log/slog"
	"net/http"
	"os"
	"time"

	"github.com/saugat1070/olx-api/internal/config"
	"github.com/saugat1070/olx-api/internal/db"
	"github.com/saugat1070/olx-api/internal/handlers"
)

func main() {
	cfg := config.MustLoadConfig()
	db, errordb := db.Connect(cfg.DB_URL)
	if errordb != nil {
		log.Fatalf("main.db.connect: %v", errordb)
	}

	lh := handlers.NewListingHandler(db)

	// logger setup

	loggerHandler := slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		Level:     slog.LevelInfo,
		AddSource: true,
	})

	logger := slog.New(loggerHandler)
	slog.SetDefault(logger)

	mux := http.NewServeMux()
	mux.HandleFunc("GET /healthz", handlers.Health)
	mux.HandleFunc("GET /listing", lh.Listing)
	mux.HandleFunc("DELETE /listing/{id}", lh.RemoveListing)

	// http server struct
	server := http.Server{
		Addr:         ":" + cfg.PORT,
		Handler:      mux,
		ReadTimeout:  time.Second * 10,
		WriteTimeout: time.Second * 30,
		IdleTimeout:  time.Second * 60,
	}
	log.Printf("server is listening on %s", cfg.PORT)
	err := server.ListenAndServe()

	if err != nil {
		panic(err)
	}

}
