package routes

import (
	"database/sql"

	"github.com/gorilla/mux"
	"GeoTagger/controllers"
)

func SetupRoutes(db *sql.DB) *mux.Router {
	r := mux.NewRouter()

	// Registration route
	r.HandleFunc("/api/register", controllers.RegisterUser(db)).Methods("POST")

	// Login route
	r.HandleFunc("/api/login", controllers.Login(db)).Methods("POST")

	// Add more routes here as needed (e.g., for notes, tags, etc.)

	return r
}
