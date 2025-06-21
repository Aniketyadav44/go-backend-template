package config

import (
	"database/sql"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

type Config struct {
	Port string
	DB   *sql.DB
}

func LoadConfig() (*Config, error) {
	if err := godotenv.Load(); err != nil {
		return nil, err
	}

	port := getEnvKey("PORT", "")
	dbHost := getEnvKey("DB_HOST", "")
	dbPort := getEnvKey("DB_PORT", "")
	dbUser := getEnvKey("DB_USERNAME", "")
	dbPassword := getEnvKey("DB_PASSWORD", "")
	dbName := getEnvKey("DB_NAME", "")

	db, err := loadDb(dbHost, dbPort, dbUser, dbPassword, dbName)
	if err != nil {
		return nil, err
	}

	return &Config{
		Port: port,
		DB:   db,
	}, nil
}

func getEnvKey(key, defaultValue string) string {
	if val, exists := os.LookupEnv(key); exists {
		return val
	}
	return defaultValue
}
