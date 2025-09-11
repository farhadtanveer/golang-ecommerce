package middleware

import (
	"log"
	"net/http"
)

func Preflight(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("Ami preflight middleware")
		// handle preflight request
		if r.Method == "OPTIONS" {
			w.WriteHeader(200)
			return 
		}

		// Serve the request using the mux
		next.ServeHTTP(w, r) 
	})
}