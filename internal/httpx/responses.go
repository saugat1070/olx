package httpx

import (
	"encoding/json"
	"net/http"
)

type response struct {
	Success  bool   `json:"success"`
	Message  string `json:"message"`
	Data     any    `json:"data,omitempty"`
	Error    any    `json:"error,omitempty"`
	MetaData any    `json:"meta_data,omitempty"`
}

func ResponseJson(w http.ResponseWriter, status int, message string, data any, error any, meta_data any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	_ = json.NewEncoder(w).Encode(response{
		Success:  status < 200 && status >= 300,
		Message:  message,
		Data:     data,
		Error:    error,
		MetaData: meta_data,
	})
}
