package routes

import (
	"database/sql"
	"GeoTagger/middlewares"
	"github.com/gorilla/mux"
	"GeoTagger/controllers"
	
)

func SetupRoutes(db *sql.DB) *mux.Router {
	r := mux.NewRouter()

	r.Use(middlewares.CORSMiddleware) // Apply middleware globally

	// Auth
	r.HandleFunc("/api/register", controllers.RegisterUser(db)).Methods("POST")
	r.HandleFunc("/api/login", controllers.Login(db)).Methods("POST")
	r.HandleFunc("/api/verify", controllers.VerifyTokenHandler()).Methods("GET")

	// Notes
	r.HandleFunc("/api/notes", controllers.CreateNote(db)).Methods("POST")
	r.HandleFunc("/api/notes", controllers.GetAllNotes(db)).Methods("GET")
	r.HandleFunc("/api/notes/nearby", controllers.GetNearbyNotes(db)).Methods("GET")
	r.HandleFunc("/api/notes/{id}", controllers.UpdateNote(db)).Methods("PUT")
	r.HandleFunc("/api/notes/{id}", controllers.DeleteNote(db)).Methods("DELETE")

	return r
}

