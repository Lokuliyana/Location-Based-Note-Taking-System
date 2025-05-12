package routes

import (
	"database/sql"

	"github.com/gorilla/mux"
	"GeoTagger/controllers"
)

func SetupRoutes(db *sql.DB) *mux.Router {
	r := mux.NewRouter()

	
	r.HandleFunc("/api/register", controllers.RegisterUser(db)).Methods("POST")

	r.HandleFunc("/api/login", controllers.Login(db)).Methods("POST")

	r.HandleFunc("/api/notes", controllers.CreateNote(db)).Methods("POST")

	r.HandleFunc("/api/notes/nearby", controllers.GetNearbyNotes(db)).Methods("GET")

	r.HandleFunc("/api/notes", controllers.GetAllNotes(db)).Methods("GET")

	r.HandleFunc("/api/notes", controllers.UpdateNote(db)).Methods("PUT")

	r.HandleFunc("/api/notes", controllers.DeleteNote(db)).Methods("DELETE")

	r.HandleFunc("/api/verify", controllers.VerifyTokenHandler()).Methods("GET")



	// Add more routes here as needed (e.g., for notes, tags, etc.)

	return r
}
