package main

import (
    "log"
    "database/sql"
    "GeoTagger/config"
    "GeoTagger/routes"
    "net/http"  // Ensure to import net/http package for HTTP server
    _ "github.com/go-sql-driver/mysql" // MySQL driver
)

func main() {
    // Load environment variables from .env file
    config.LoadEnv()

    
    dbUser := config.GetEnv("DB_USER")
    dbPassword := config.GetEnv("DB_PASSWORD")
    dbHost := config.GetEnv("DB_HOST")
    dbPort := config.GetEnv("DB_PORT")
    dbName := config.GetEnv("DB_NAME")

    // Construct the Data Source Name (DSN) for connecting to the MySQL database
    dsn := dbUser + ":" + dbPassword + "@tcp(" + dbHost + ":" + dbPort + ")/" + dbName

    // Open a connection to the database
    db, err := sql.Open("mysql", dsn)
    if err != nil {
        log.Fatalf("Error connecting to the database: %v", err)
    }
    defer db.Close() // Ensure to close the DB connection when the function returns

    // Set up the routes for the application
    r := routes.SetupRoutes(db)

    // Start the HTTP server on port 5000
    log.Println("Server started on :5000")
    log.Fatal(http.ListenAndServe(":5000", r))  // Start the server and handle errors
}
