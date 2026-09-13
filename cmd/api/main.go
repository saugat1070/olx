package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		log.Fatalf("Error occured while running server: %s", err)
	}
	fmt.Println("Server is running on port 8000")
}
