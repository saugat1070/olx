package handlers

import (
	"database/sql"
	"encoding/json"
	"log/slog"
	"net/http"

	"github.com/saugat1070/olx-api/internal/httpx"
	"golang.org/x/crypto/bcrypt"
)

type Authentication struct {
	db  *sql.DB
	log *slog.Logger
}

func NewAuthenticationHandler(db *sql.DB, log *slog.Logger) *Authentication {
	return &Authentication{
		db:  db,
		log: log,
	}
}

func (auth *Authentication) CreateUser(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var requestBody SignupRequest
	if err := json.NewDecoder(r.Body).Decode(&requestBody); err != nil {
		auth.log.Error("failed to decode", "error:", err.Error())
		http.Error(w, "internal Server Error", http.StatusInternalServerError)
		return
	}

	if err := requestBody.Validate(); err != nil {
		auth.log.Error("failed to validate", "error:", err.Error())
		http.Error(w, "internal Server Error", http.StatusInternalServerError)
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(requestBody.Password), 10)
	if err != nil {
		auth.log.Error("failed to hash the password", "error: ", err.Error())
		return
	}
	sqlQuery := `INSERT INTO users (email,password,full_name) VALUES ($1,$2,$3) RETURNING email,full_name `
	row := auth.db.QueryRowContext(ctx, sqlQuery, requestBody.Email, hashedPassword, requestBody.FullName)

	if err := row.Scan(); err != nil {
		auth.log.Error("failed to create user", "error:", err.Error())
		http.Error(w, "internal Server Error", http.StatusInternalServerError)
		return
	}

	httpx.ResponseJson(w, http.StatusCreated, "user create successfully", &SignupResponse{
		Email:    requestBody.Email,
		FullName: requestBody.FullName,
	}, nil, nil)

}
