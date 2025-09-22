package rest

import (
	"ecommerce/config"
	"ecommerce/rest/handlers/product"
	"ecommerce/rest/handlers/user"
	middleware "ecommerce/rest/middleswares"
	"fmt"
	"net/http"
	"os"
	"strconv"
)

type Server struct {
	cnf 		config.Config
	productHandler *product.Handler
	userHandler    *user.Handler
}

func NewServer(
	cnf config.Config,
	productHandler *product.Handler, 
	userHandler *user.Handler,
) *Server {
	return &Server{
		cnf:            cnf,
		productHandler: productHandler,
		userHandler:    userHandler,
	}
}

func (server *Server) Start() {
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
	// initRoutes(mux, manager)
	server.productHandler.RegisterRoutes(mux, manager)
	server.userHandler.RegisterRoutes(mux, manager)

	
	addr := ":" + strconv.Itoa(server.cnf.HttpPort)
	fmt.Println("Server is running on port", addr)
	err := http.ListenAndServe(addr, wrappedMux) // failed to start the server
	if err != nil {
		fmt.Println("Error starting server:", err)
		os.Exit(1)
	}
}