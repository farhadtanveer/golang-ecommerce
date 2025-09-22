package cmd

import (
	"ecommerce/config"
	"ecommerce/rest"
)

func Serve() {
	cnf := config.GetConfig() // Load configuration
	rest.Start(cnf)       // Start the server
}