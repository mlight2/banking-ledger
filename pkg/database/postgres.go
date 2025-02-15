package database

import (
	"fmt"
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

// Config structure to hold database configuration
type Config struct {
	PostgresURL string
}

var DB *sqlx.DB

// InitPostgres initializes PostgreSQL connection
func InitPostgres(cfg *Config) {
	var err error
	DB, err = sqlx.Connect("postgres", cfg.PostgresURL)
	if err != nil {
		log.Fatal("Error connecting to PostgreSQL:", err)
	}

	fmt.Println("Connected to PostgreSQL successfully!")
}
