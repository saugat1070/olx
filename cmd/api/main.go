package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {

	http.HandleFunc("GET /health", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("okay")) // conversion from string to byte
	})

	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		log.Fatalf("Error occured while running server: %s", err)
	}
	fmt.Println("Server is running on port 8000")
}
