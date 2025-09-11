package cmd

import (
	"ecommerce/config"
	"ecommerce/middleware"
	"fmt"
	"net/http"
	"os"
	"strconv"
)

func Serve() {
	cnf := config.GetConfig() // Load configuration

	// Initialize the HTTP server and middleware manager
	manager := middleware.NewManager()
	manager.Use(
		middleware.Cors, 
		middleware.Preflight, 
		middleware.Logger,
	)

	mux := http.NewServeMux()
	wrappedMux := manager.WrapMux(mux)

	// Initialize routes
	initRoutes(mux, manager)

	
	addr := ":" + strconv.Itoa(cnf.HttpPort)
	fmt.Println("Server is running on port", addr)
	err := http.ListenAndServe(addr, wrappedMux) // failed to start the server
	if err != nil {
		fmt.Println("Error starting server:", err)
		os.Exit(1)
	}
}