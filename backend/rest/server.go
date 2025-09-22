package rest

import (
	"ecommerce/config"
	middleware "ecommerce/rest/middleswares"
	"fmt"
	"net/http"
	"os"
	"strconv"
)

func Start(cnf config.Config) {
	// Initialize the HTTP server and middleware manager
	manager := middleware.NewManager()
	manager.Use(
		middleware.Preflight,
		middleware.Cors, 
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