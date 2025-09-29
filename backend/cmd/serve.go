package cmd

import (
	"ecommerce/config"
	"ecommerce/infra/db"
	"ecommerce/repo"
	"ecommerce/rest"
	"ecommerce/rest/handlers/product"
	"ecommerce/rest/handlers/user"
	middleware "ecommerce/rest/middleswares"
	"fmt"
)

func Serve() {
	cnf := config.GetConfig()// Load configuration
	dbCon, err := db.NewConnection() // Initialize database connection
	if err != nil {
		fmt.Println("Failed to connect to the database:", err)
	}

	middlewares := middleware.NewMiddlewares(cnf) // Initialize middlewares

	// Initialize repositories
	productRepo := repo.NewProductRepo() 
	userRepo := repo.NewUserRepo(dbCon)

	// Initialize handlers
	productHandler := product.NewHandler(middlewares, productRepo)
	userHandler := user.NewHandler(cnf, userRepo)

	// Start the server
	server := rest.NewServer(cnf, productHandler, userHandler)      
	server.Start()
}