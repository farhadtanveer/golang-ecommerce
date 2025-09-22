package cmd

import (
	"ecommerce/config"
	"ecommerce/rest"
	"ecommerce/rest/handlers/product"
	"ecommerce/rest/handlers/user"
)

func Serve() {
	cnf := config.GetConfig()// Load configuration

	productHandler := product.Newhandler() // Initialize handlers
	userHandler := user.Newhandler()

	server := rest.NewServer(cnf, productHandler, userHandler)      // Start the server
	server.Start()
}