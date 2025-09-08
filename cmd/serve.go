package cmd

import (
	"ecommerce/global_router"
	"ecommerce/middleware"
	"fmt"
	"net/http"
)

func Serve() {

	// Initialize the HTTP server and middleware manager
	mux := http.NewServeMux()
	manager := middleware.NewManager()
	manager.Use( middleware.Logger,middleware.Hudai) // Apply global middleware
	
	// Initialize routes
	initRoutes(mux, manager)

	fmt.Println("Server is running on port 8080")

	globalRouter := global_router.GlobalRouter(mux)

	err := http.ListenAndServe(":8080", globalRouter) // failed to start the server
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}