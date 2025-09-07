package cmd

import (
	"ecommerce/global_router"
	"ecommerce/handlers"
	"ecommerce/middleware"
	"fmt"
	"net/http"
)

func Serve() {
	mux := http.NewServeMux()


	mux.Handle("GET /route", middleware.Logger(http.HandlerFunc(handlers.Test)))
	mux.Handle("GET /products", http.HandlerFunc(handlers.CreateProduct))
	mux.HandleFunc("POST /products", http.HandlerFunc(handlers.CreateProduct))
	mux.HandleFunc("GET /products/{id}", http.HandlerFunc(handlers.GetProductByID))

	fmt.Println("Server is running on port 8080")

	globalRouter := global_router.GlobalRouter(mux)

	err := http.ListenAndServe(":8080", globalRouter) // failed to start the server
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}