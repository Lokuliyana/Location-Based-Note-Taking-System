package controllers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

func GetNearbyNotes(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Parse latitude, longitude, and radius from query parameters
		latStr := r.URL.Query().Get("latitude")
		lonStr := r.URL.Query().Get("longitude")
		radiusStr := r.URL.Query().Get("radius")

		if latStr == "" || lonStr == "" || radiusStr == "" {
			http.Error(w, "Missing query parameters", http.StatusBadRequest)
			return
		}

		latitude, err1 := strconv.ParseFloat(latStr, 64)
		longitude, err2 := strconv.ParseFloat(lonStr, 64)
		radius, err3 := strconv.ParseFloat(radiusStr, 64) // in meters

		if err1 != nil || err2 != nil || err3 != nil {
			http.Error(w, "Invalid latitude, longitude, or radius", http.StatusBadRequest)
			return
		}

		// SQL query to find notes within radius using ST_Distance_Sphere
		query := `
			SELECT id, user_id, title, content, 
				   ST_X(location) AS longitude, ST_Y(location) AS latitude, 
				   created_at 
			FROM notes
			WHERE ST_Distance_Sphere(location, ST_GeomFromText(?)) <= ?
		`

		point := fmt.Sprintf("POINT(%f %f)", longitude, latitude)

		rows, err := db.Query(query, point, radius)
		if err != nil {
			http.Error(w, "Database query error", http.StatusInternalServerError)
			return
		}
		defer rows.Close()

		type Note struct {
			ID        int     `json:"id"`
			UserID    int     `json:"user_id"`
			Title     string  `json:"title"`
			Content   string  `json:"content"`
			Longitude float64 `json:"longitude"`
			Latitude  float64 `json:"latitude"`
			CreatedAt string  `json:"created_at"`
		}

		var notes []Note
		for rows.Next() {
			var note Note
			if err := rows.Scan(&note.ID, &note.UserID, &note.Title, &note.Content, &note.Longitude, &note.Latitude, &note.CreatedAt); err != nil {
				http.Error(w, "Failed to read results", http.StatusInternalServerError)
				return
			}
			notes = append(notes, note)
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(notes)
	}
}
