package global_router

import "net/http"

func GlobalRouter(mux *http.ServeMux) http.Handler {
	handleAllReq := func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, DELETE, PUT, PATCH, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		w.Header().Set("content-Type", "application/json")
		if r.Method == "OPTIONS" {
			w.WriteHeader(200)// Handle preflight request
			return 
		}
		mux.ServeHTTP(w, r) // Serve the request using the mux
	}

	return http.HandlerFunc(handleAllReq)
}