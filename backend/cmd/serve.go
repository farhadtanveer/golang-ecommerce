package cmd

import (
	"ecommerce/config"
	"ecommerce/infra/db"
	"ecommerce/repo"
	"ecommerce/rest"
	prdctHandler "ecommerce/rest/handlers/product"
	usrHandler "ecommerce/rest/handlers/user"
	middleware "ecommerce/rest/middleswares"
	"ecommerce/user"

	"fmt"
	"os"
)

func Serve() {
	cnf := config.GetConfig()// Load configuration
	dbCon, err := db.NewConnection(cnf.DB) // Initialize database connection
	if err != nil {
		fmt.Println("Failed to connect to the database:", err)
	}

	middlewares := middleware.NewMiddlewares(cnf) // Initialize middlewares

	// Initialize repositories / repos
	productRepo := repo.NewProductRepo(dbCon) 
	userRepo := repo.NewUserRepo(dbCon)

	// domains
	usrSvc := user.NewService(userRepo)

	err = db.MigrateDB(dbCon, "./migrations") // Run database migrations
	if err != nil {
		fmt.Println("Failed to run migrations:", err)
		os.Exit(1)
	}

	// Initialize handlers
	productHandler := prdctHandler.NewHandler(middlewares, productRepo)
	userHandler := usrHandler.NewHandler(cnf, usrSvc)

	// Start the server
	server := rest.NewServer(cnf, productHandler, userHandler)      
	server.Start()
}