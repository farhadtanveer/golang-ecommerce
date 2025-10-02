package db

import (
	"ecommerce/config"
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func GetConnectionString(cnf *config.DBConfig) string {
	connString := fmt.Sprintf("user=%s password=%s host=%s port=%d dbname=%s sslmode=%s",
		cnf.User,
		cnf.Password,
		cnf.Host,
		cnf.Port,
		cnf.Name,
	)
	if !cnf.SSLMode {
		connString += " sslmode=disable"
	}

	return connString
}

func NewConnection(cnf *config.DBConfig) (*sqlx.DB, error) {
	dbsource := GetConnectionString(cnf)
	dbCon, err := sqlx.Connect("postgres", dbsource)
	if err != nil {
		return nil, err
	}
	return dbCon, nil
}