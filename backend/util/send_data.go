package util

import (
	"encoding/json"
	"net/http"
)

func SendData(w http.ResponseWriter, statusCode int, data interface{}) {
	w.Header().Set("Content-Type", "application/json") // ✅ Important
	w.WriteHeader(statusCode)
	encoder := json.NewEncoder(w)
	encoder.Encode(data)
}

func SendError(w http.ResponseWriter, statusCode int, msg string) {
	w.Header().Set("Content-Type", "application/json") // ✅ Important
	w.WriteHeader(statusCode)
	encoder := json.NewEncoder(w)
	encoder.Encode(map[string]string{"error": msg}) // better format for error
}
