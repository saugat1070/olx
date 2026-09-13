package handlers

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"time"
)

type listing struct {
	ID          string    `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Price       int64     `json:"price"`
	City        string    `json:"city"`
	CreatedAt   time.Time `json:"created_at"`
}

func Listing(db *sql.DB) http.HandlerFunc { // Closure factory
	return func(w http.ResponseWriter, r *http.Request) {
		rows, err := db.Query(
			`
			SELECT * FROM listings
			ORDER BY created_at DESC
			LIMIT 100
			`)
		if err != nil {
			log.Printf("Query: %v", err)
			http.Error(w, "internal Server Error", http.StatusInternalServerError)
			return
		}
		defer rows.Close()
		listings := []listing{}
		for rows.Next() {
			var list listing
			err := rows.Scan(&list.ID, &list.Title, &list.Description, &list.Price, &list.City, &list.CreatedAt)
			if err != nil {
				log.Printf("rows.Scan: %v", err)
				http.Error(w, "internal Server Error", http.StatusInternalServerError)
				return
			}
			listings = append(listings, list)
		}
		if err := rows.Err(); err != nil {
			log.Printf("rows.Err: %v", err)
			http.Error(w, "internal Server Error", http.StatusInternalServerError)
			return
		}

		w.Header().Set("content-type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(listings)

	}
}
