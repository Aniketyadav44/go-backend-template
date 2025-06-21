package config

import (
	"database/sql"
	"fmt"
)

func loadDb(dbHost, dbPort, dbUser, dbPassword, dbName string) (*sql.DB, error) {
	psConnStr := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", dbUser, dbPassword, dbHost, dbPort, dbName)

	db, err := sql.Open("postgres", psConnStr)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
