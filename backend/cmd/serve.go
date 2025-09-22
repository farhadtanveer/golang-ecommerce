package cmd

import (
	"ecommerce/config"
	"ecommerce/rest"
	"ecommerce/rest/handlers/product"
	"ecommerce/rest/handlers/user"
	middleware "ecommerce/rest/middleswares"
)

func Serve() {
	cnf := config.GetConfig()// Load configuration

	middlewares := middleware.NewMiddlewares(cnf) // Initialize middlewares

	productHandler := product.NewHandler(middlewares) // Initialize handlers
	userHandler := user.NewHandler()

	server := rest.NewServer(cnf, productHandler, userHandler)      // Start the server
	server.Start()
}