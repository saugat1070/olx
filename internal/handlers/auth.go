package handlers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"

	"github.com/saugat1070/olx-api/internal/config"
	"github.com/saugat1070/olx-api/internal/helper"
	"github.com/saugat1070/olx-api/internal/httpx"
	"golang.org/x/crypto/bcrypt"
)

type Authentication struct {
	cfg config.Config
	db  *sql.DB
	log *slog.Logger
}

func NewAuthenticationHandler(cfg config.Config, db *sql.DB, log *slog.Logger) *Authentication {
	return &Authentication{
		cfg: cfg,
		db:  db,
		log: log,
	}
}

func (auth *Authentication) CreateUser(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var requestBody SignupRequest
	if err := json.NewDecoder(r.Body).Decode(&requestBody); err != nil {
		auth.log.Error("failed to decode", "error:", err.Error())
		httpx.Error(w, http.StatusBadRequest, "", httpx.CodeMalformedJSON, err.Error())
		return
	}

	if err := requestBody.Validate(); err != nil {
		auth.log.Error("failed to validate", "error:", err.Error())
		httpx.Error(w, http.StatusInternalServerError, "internal Server Error", httpx.CodeValidationFailed, err.Error())
		return
	}

	// check if user is already exists
	userExistsQuery := `SELECT id FROM users WHERE email=$1`
	var user string
	selectUserRow := auth.db.QueryRowContext(ctx, userExistsQuery, requestBody.Email)
	if err := selectUserRow.Scan(&user); err != nil {
		auth.log.Error("failed to create user", "error:", err.Error())
		httpx.Error(w, http.StatusInternalServerError, "", httpx.CodeConflict, "email")
		return
	}
	fmt.Println(user)
	if user != "" {
		httpx.ResponseJson(w, http.StatusAccepted, fmt.Sprintf("User already exists with this email: %s", requestBody.Email), nil, nil, nil)
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(requestBody.Password), 10)
	if err != nil {
		auth.log.Error("failed to hash the password", "error: ", err.Error())
		return
	}
	sqlQuery := `INSERT INTO users (email,password,full_name) VALUES ($1,$2,$3) `
	row := auth.db.QueryRowContext(ctx, sqlQuery, requestBody.Email, hashedPassword, requestBody.FullName)

	if err := row.Scan(); err != nil {
		auth.log.Error("failed to create user", "error:", err.Error())
		httpx.Error(w, http.StatusInternalServerError, "internal Server Error", httpx.CodeConflict, "email")
		return
	}

	httpx.ResponseJson(w, http.StatusCreated, "user create successfully", &SignupResponse{
		Email:    requestBody.Email,
		FullName: requestBody.FullName,
	}, nil, nil)

}

func (auth *Authentication) LoginUser(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var LoginRequest LoginRequest
	if err := json.NewDecoder(r.Body).Decode(&LoginRequest); err != nil {
		auth.log.Error("failed to decode", "error:", err.Error())
		http.Error(w, "internal Server Error", http.StatusInternalServerError)
		return
	}

	var user User
	sqlQuery := `SELECT id,email,password,full_name,role FROM users WHERE email = $1`
	row := auth.db.QueryRowContext(ctx, sqlQuery, LoginRequest.Email)
	if err := row.Scan(&user.ID, &user.Email, &user.Password, &user.FullName, &user.Role); err != nil {
		auth.log.Error("failed to find user", "error:", err.Error())
		http.Error(w, "internal Server Error", http.StatusInternalServerError)
		return
	}

	if user.Email == "" {
		httpx.Error(w, http.StatusUnauthorized, fmt.Sprintf("user with email: %s is not exists", LoginRequest.Email), httpx.CodeUnauthenticated, "")
		return
	}

	accessTokenPayload := helper.JwtPayload{
		ID:       user.ID,
		Email:    user.Email,
		FullName: user.FullName,
		Role:     user.Role,
	}

	accessToken, err := helper.CreateAccessToken(accessTokenPayload, auth.cfg.JWT_SECRET_KEY)
	if err != nil {
		auth.log.Error("failed to create access token", "error:", err.Error())
		http.Error(w, "internal Server Error", http.StatusInternalServerError)
		return
	}

	httpx.ResponseJson(w, http.StatusOK, "user login successfully", &LoginResponse{
		AccessToken: accessToken,
	}, nil, nil)

}
