package config

import (
	"database/sql"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/go-sql-driver/mysql"
)

// LoadEnv loads environment variables from a .env file if present
func LoadEnv() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using environment variables")
	}
}

// GetEnv retrieves the value of an environment variable, or returns a default value if not found
func GetEnv(key string) string {
	value := os.Getenv(key)
	if value == "" {
		log.Fatalf("Environment variable %s not set", key)
	}
	return value
}

// ConnectDB establishes a connection to the MySQL database using credentials from environment variables
func ConnectDB() *sql.DB {
	// Build the DSN from environment variables
	dbUser := GetEnv("DB_USER")
	dbPassword := GetEnv("DB_PASSWORD")
	dbHost := GetEnv("DB_HOST")
	dbPort := GetEnv("DB_PORT")
	dbName := GetEnv("DB_NAME")

	// Construct the Data Source Name (DSN) string for MySQL connection
	dsn := dbUser + ":" + dbPassword + "@tcp(" + dbHost + ":" + dbPort + ")/" + dbName

	// Open a connection to the MySQL database
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal("Error connecting to DB:", err)
	}

	// Test the connection
	if err = db.Ping(); err != nil {
		log.Fatal("DB unreachable:", err)
	}

	return db
}
