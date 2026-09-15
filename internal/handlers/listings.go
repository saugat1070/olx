package handlers

import (
	"database/sql"
	"encoding/json"
	"errors"
	"log/slog"
	"net/http"
	"time"

	"github.com/saugat1070/olx-api/internal/httpx"
	"github.com/saugat1070/olx-api/internal/middleware"
)

type listing struct {
	ID          string    `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Price       int64     `json:"price"`
	City        string    `json:"city"`
	CreatedAt   time.Time `json:"created_at"`
}

type ListingHandler struct {
	db     *sql.DB
	logger *slog.Logger
}

// constructor function for Listinghandler
func NewListingHandler(db *sql.DB, logger *slog.Logger) *ListingHandler {
	return &ListingHandler{
		db:     db,
		logger: logger,
	}
}

// lh -> method receiver from listinghandler struct
func (lh *ListingHandler) Listing(w http.ResponseWriter, r *http.Request) {
	// Request context
	ctx := r.Context()
	rows, err := lh.db.QueryContext(ctx,
		`
			SELECT id, title, description, price, city, created_at FROM listings
			ORDER BY created_at DESC
			LIMIT 100
			`)
	if err != nil {
		lh.logger.Error("Error in sql.Query in list handler", "error: ", err.Error())
		http.Error(w, "internal Server Error", http.StatusInternalServerError)
		return
	}
	defer rows.Close()
	listings := []listing{}
	for rows.Next() {
		var list listing
		err := rows.Scan(&list.ID, &list.Title, &list.Description, &list.Price, &list.City, &list.CreatedAt)
		if err != nil {
			lh.logger.Error("Error in sql.Rows in list handler", "error: ", err.Error())
			http.Error(w, "internal Server Error", http.StatusInternalServerError)
			return
		}
		listings = append(listings, list)
	}
	if err := rows.Err(); err != nil {
		lh.logger.Error("Error in sql.Rows in list handler", "error: ", err.Error())
		http.Error(w, "internal Server Error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(listings)
}

func (lh *ListingHandler) RemoveListing(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	id := r.PathValue("id")

	_, err := lh.db.ExecContext(ctx,
		`
			DELETE FROM list WHERE id = $1
			`,
		id)

	if err != nil {
		lh.logger.Error("Error on executing delete opertion in listinghandler", "listing_id: ", id, "request-id: ", ctx.Value("X-Request-ID"), "error:", err.Error())
		http.Error(w, "internal Server Error", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
	json.NewEncoder(w).Encode(`{"status:"list delete successfully""}`)
}

func (lh *ListingHandler) CreateList(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	requestId := middleware.GetRequestID(ctx)
	var req CreateListingRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		lh.logger.Error("failed to decode", "request_id: ", requestId, "error: ", err.Error())
		httpx.Error(w, http.StatusBadRequest, "please provide proper listing", httpx.CodeMalformedJSON, "")
		return
	}

	if err := req.Validate(); err != nil {
		var varr *ValidationError
		errors.As(err, &varr) // it will check if the error is of the type ValidationError and if it is it will extract the value
		httpx.Error(w, http.StatusUnprocessableEntity, err.Error(), httpx.CodeValidationFailed, varr.Field)
		return
	}

	row := lh.db.QueryRowContext(ctx,
		`
		INSERT INTO listings (title,description,price,city) 
		VALUES ($1,$2,$3,$4)
		RETURNING id
	`,
		req.Title, req.Description, req.Price, req.City)

	if err := row.Scan(); err != nil {
		lh.logger.Error("failed to insert", "request_id: ", requestId, "error: ", err.Error())
		http.Error(w, "failed to create listing", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]any{
		"status":  "success",
		"message": "listing created successfully",
		"data":    req,
	})
}
