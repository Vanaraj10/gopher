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
        log.Println("Error loading .env file, using environment variables directly")
    }
    dbURL := os.Getenv("DATABASE_URL")

    var errOpen error
    DB, errOpen = sql.Open("postgres", dbURL)
    if errOpen != nil {
        log.Fatalf("Error connecting to the database: %v", errOpen)
        return
    }
    errPing := DB.Ping()
    if errPing != nil {
        log.Fatalf("Error connecting to the database: %v", errPing)
        return
    }
    log.Println("Connected to database successfully")
}