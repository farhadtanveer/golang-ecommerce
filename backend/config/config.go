package config

import (
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var configuration *Config

type DBConfig struct {
	Host     string `env:"DB_HOST"`
	Port     int
	User     string
	Password string
	Name   string
	SSLMode  bool
}

type Config struct {
	Version     string
	ServiceName string
	HttpPort    int
	JwtSecretKey string
	DB   *DBConfig
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

	jwtSecretKey := os.Getenv("JWT_SECRET_KEY")
	if jwtSecretKey == "" {
		fmt.Println("JWT_SECRET_KEY not set, using default")
		os.Exit(1)
	}

	dbHost := os.Getenv("DB_HOST")
	if dbHost == "" {
		fmt.Println("Host is required")
		os.Exit(1)
	}

	dbPort := os.Getenv("DB_PORT")
	if dbPort == "" {
		fmt.Println("Port is required")
		os.Exit(1)
	}

	dbPortInt, err := strconv.ParseInt(dbPort, 10, 64)
	if err != nil {
		fmt.Println("Port must be number")
		os.Exit(1)
	}

	dbName := os.Getenv("DB_NAME")
	if dbName == "" {
		fmt.Println("NAME is required")
		os.Exit(1)
	}

	dbUser := os.Getenv("DB_USER")
	if dbUser == "" {
		fmt.Println("USER is required")
		os.Exit(1)
	}

	dbPassword := os.Getenv("DB_PASSWORD")
	if dbPassword == "" {
		fmt.Println("PASSWORD is required")
		os.Exit(1)
	}

	dbSSLMode := os.Getenv("DB_ENABLE_SSL_MODE")
	enableSSLMode, err := strconv.ParseBool(dbSSLMode)
	if err != nil {
		enableSSLMode = false // default is false
	}
	

	dbConfig := &DBConfig{
		Host:     dbHost,
		Port:     int(dbPortInt),
		User:     dbUser,
		Password: dbPassword,
		Name:     dbName,
		SSLMode:  enableSSLMode,
		}

	configuration = &Config{
		Version:     version,
		ServiceName: serviceName,
		HttpPort:    int(port),
		JwtSecretKey: jwtSecretKey,
		DB:          dbConfig,
	}
}

func GetConfig() *Config{
	if configuration == nil {
		loadConfig()
	}
	return configuration
}