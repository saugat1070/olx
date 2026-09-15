package handlers

import (
	"fmt"
	"strings"
)

type CreateListingRequest struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Price       int64  `json:"price"`
	City        string `json:"city"`
}

type ValidationError struct {
	Field   string `json:"field,omitempty"`
	Message string `json:"message"`
}

func (ve *ValidationError) Error() string {
	return fmt.Sprintf("validation failed on field %s: %s", ve.Field, ve.Message)
}

func (req *CreateListingRequest) Validate() error {
	if strings.TrimSpace(req.Title) == "" {
		return &ValidationError{
			Field:   "title",
			Message: "must not be empty",
		}
	}
	if len(strings.TrimSpace(req.Description)) > 500 {
		return &ValidationError{
			Field:   "description",
			Message: "must be at most 500 characters long",
		}
	}
	if req.Price <= 0 {
		return &ValidationError{
			Field:   "price",
			Message: "must be greater than 0",
		}
	}
	if strings.TrimSpace(req.City) == "" {
		return &ValidationError{
			Field:   "city",
			Message: "must not be empty",
		}
	}
	return nil
}
