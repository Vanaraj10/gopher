package config

import (
	"database/sql"
	"log"
	"os"
	_ "github.com/lib/pq" // PostgreSQL driver
	"github.com/joho/godotenv"
)

var DB *sql.DB

func ConnectDB() {

	err := godotenv.Load()
	if err != nil {
		panic("Error loading .env file")
	}
	dbURL := os.Getenv("DATABASE_URL")

	DB, err := sql.Open("postgres", dbURL)
	if err != nil {
		panic("Failed to connect to database: " + err.Error())
		
	}
	err = DB.Ping()
	if err != nil {
		panic("Failed to ping database: " + err.Error())
	}
	log.Println("Connected to database successfully")
}