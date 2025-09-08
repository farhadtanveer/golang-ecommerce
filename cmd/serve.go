package cmd

import (
	"ecommerce/middleware"
	"fmt"
	"net/http"
)

func Serve() {

	// Initialize the HTTP server and middleware manager
	manager := middleware.NewManager()

	mux := http.NewServeMux()
	
	wrappedMux := manager.WrapMux(
		mux, 
		middleware.Logger,
		middleware.Hudai, 
		middleware.CorsWithPreflight,
	) // Wrap the mux with middleware

	// Initialize routes
	initRoutes(mux, manager)

	fmt.Println("Server is running on port 8080")

	err := http.ListenAndServe(":8080", wrappedMux) // failed to start the server
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}