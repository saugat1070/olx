package handlers

import (
	"regexp"
	"strings"
)

type SignupRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	FullName string `json:"full_name"`
}

type SignupResponse struct {
	Email    string `json:"email"`
	FullName string `json:"full_name"`
}

func (req *SignupRequest) Validate() error {
	regex := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	if strings.TrimSpace(req.Email) == "" && !regex.MatchString(req.Email) {
		return &ValidationError{
			Field:   "email",
			Message: "must not be empty and must be a valid email address",
		}
	}
	if len(strings.TrimSpace(req.Password)) < 8 {
		return &ValidationError{
			Field:   "password",
			Message: "must be at least 8 characters long",
		}
	}

	return nil
}
