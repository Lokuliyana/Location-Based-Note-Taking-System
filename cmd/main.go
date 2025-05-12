package main

import (
    "database/sql"
    "log"
    "net/http"

    "github.com/rs/cors"
    _ "github.com/go-sql-driver/mysql"

    "GeoTagger/config"
    "GeoTagger/routes"
)

func main() {
    // Load environment variables from .env file
    config.LoadEnv()

    dbUser := config.GetEnv("DB_USER")
    dbPassword := config.GetEnv("DB_PASSWORD")
    dbHost := config.GetEnv("DB_HOST")
    dbPort := config.GetEnv("DB_PORT")
    dbName := config.GetEnv("DB_NAME")

    // Construct the DSN (Data Source Name)
    dsn := dbUser + ":" + dbPassword + "@tcp(" + dbHost + ":" + dbPort + ")/" + dbName

    // Open the database connection
    db, err := sql.Open("mysql", dsn)
    if err != nil {
        log.Fatalf("Error connecting to the database: %v", err)
    }
    defer db.Close()

    // Set up routes with database
    router := routes.SetupRoutes(db)

    // Set up CORS middleware
    c := cors.New(cors.Options{
        AllowedOrigins:   []string{"http://localhost:5173"},
        AllowedMethods:   []string{"GET", "POST", "OPTIONS"},
        AllowedHeaders:   []string{"Content-Type", "Authorization"},
        AllowCredentials: true,
    })

    // Apply CORS to the router
    handler := c.Handler(router)

    log.Println("Server started on :5000")
    log.Fatal(http.ListenAndServe(":5000", handler))
}
