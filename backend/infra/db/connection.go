package db

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func GetConnectionString() string {
	return "user=postgres password=12345678 host=localhost port=5432 dbname=ecommerce sslmode=disable" 
}

func NewConnection() (*sqlx.DB, error) {
	dbsource := GetConnectionString()
	dbCon, err := sqlx.Connect("postgres", dbsource)
	if err != nil {
		return nil, err
	}
	return dbCon, nil
}