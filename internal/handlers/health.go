package handlers

import (
	"net/http"
)

func Health(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK) // go provide http status code as constant
	w.Header().Set("content-type", "application/json")
	w.Write([]byte(`{"status": "okay"}`)) // conversion from string to byte
}
