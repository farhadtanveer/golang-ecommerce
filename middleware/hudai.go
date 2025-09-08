package middleware

import (
	"log"
	"net/http"
)

func Hudai(next http.Handler) http.Handler{
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("Ami Hudai Middleware")
		next.ServeHTTP(w, r) // Call the next handler/controller
	})
}