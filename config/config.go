package config

import (
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var configuration Config

type Config struct {
	Version     string
	ServiceName string
	HttpPort    int
}

func loadConfig(){

	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file, trying to read from environment variables", err)
		os.Exit(1)
	}
	// connect with env or config file
	version := os.Getenv("VERSION")
	if version == "" {
		fmt.Println("VERSION not set, using default")
		os.Exit(1)
	}

	serviceName := os.Getenv("SERVICE_NAME")
	if serviceName == "" {
		fmt.Println("SERVICE_NAME not set, using default")
		os.Exit(1)
	}

	httpPort := os.Getenv("HTTP_PORT")
	if httpPort == "" {
		httpPort = "8080" // default port
	}

	port, err := strconv.ParseInt(httpPort, 10, 64) // check if port is valid number

	if err != nil {
		fmt.Println("Port must be number")
		os.Exit(1)
	}

	configuration = Config{
		Version:     version,
		ServiceName: serviceName,
		HttpPort:    int(port),
	}
}

func GetConfig() Config{
	loadConfig()
	return configuration
}