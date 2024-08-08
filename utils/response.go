package utils

import (
	"encoding/json"
	"net/http"
)

func ErrorResponse(w http.ResponseWriter, message string, code int) {
	w.WriteHeader(code)
	response := map[string]string{"error": message}
	json.NewEncoder(w).Encode(response)
}

func SuccessResponse(w http.ResponseWriter, message string, code int) {
	w.WriteHeader(code)
	response := map[string]string{"message": message}
	json.NewEncoder(w).Encode(response)
}
