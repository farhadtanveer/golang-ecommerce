package middleware

import "net/http"

func CorsWithPreflight(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// handle CORS
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, DELETE, PUT, PATCH, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		w.Header().Set("content-Type", "application/json")

		// handle preflight request
		if r.Method == "OPTIONS" {
			w.WriteHeader(200)
			return 
		}

		// Serve the request using the mux
		next.ServeHTTP(w, r) 
	})
}