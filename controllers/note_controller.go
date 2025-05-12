package controllers

import (
	"GeoTagger/utils"
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"
)

type CreateNoteRequest struct {
	Title     string  `json:"title"`
	Content   string  `json:"content"`
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

func CreateNote(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Get JWT token from Authorization header
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			http.Error(w, "Missing Authorization header", http.StatusUnauthorized)
			return
		}

		// Extract the JWT token
		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		// Parse the JWT token
		_, claims, err := utils.ParseJWT(tokenString)
		if err != nil {
			http.Error(w, "Invalid token", http.StatusUnauthorized)
			return
		}

		// Extract the user_id from the token claims
		userIDFloat, ok := claims["user_id"].(float64)
		if !ok {
			http.Error(w, "Invalid user ID in token", http.StatusUnauthorized)
			return
		}
		userID := int(userIDFloat)

		// Decode the note request from the body
		var req CreateNoteRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, "Invalid request payload", http.StatusBadRequest)
			return
		}

		// Prepare the POINT value for the location field in MySQL
		location := fmt.Sprintf("POINT(%f %f)", req.Longitude, req.Latitude)

		// Log the request data for debugging
		log.Printf("Creating note: UserID=%d, Title=%s, Latitude=%f, Longitude=%f", userID, req.Title, req.Latitude, req.Longitude)

		// Insert the note into the database
		query := `INSERT INTO notes (user_id, title, content, location, created_at)
				  VALUES (?, ?, ?, ST_GeomFromText(?), ?)`
		_, err = db.Exec(query, userID, req.Title, req.Content, location, time.Now())
		if err != nil {
			// Log the detailed error message to understand the issue
			log.Printf("Error inserting note: %v", err)
			http.Error(w, "Failed to insert note", http.StatusInternalServerError)
			return
		}

		// If everything is successful, send the success response
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{
			"message": "Note created successfully",
		})
	}
}

func GetAllNotes(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		rows, err := db.Query(`
			SELECT id, user_id, title, content, ST_X(location), ST_Y(location), created_at
			FROM notes ORDER BY created_at DESC
		`)
		if err != nil {
			http.Error(w, "Failed to fetch notes", http.StatusInternalServerError)
			return
		}
		defer rows.Close()

		type Note struct {
			ID        int     `json:"id"`
			UserID    int     `json:"user_id"`
			Title     string  `json:"title"`
			Content   string  `json:"description"`
			Longitude float64 `json:"lng"`
			Latitude  float64 `json:"lat"`
			CreatedAt string  `json:"created_at"`
		}

		var notes []Note
		for rows.Next() {
			var n Note
			err := rows.Scan(&n.ID, &n.UserID, &n.Title, &n.Content, &n.Longitude, &n.Latitude, &n.CreatedAt)
			if err != nil {
				http.Error(w, "Failed to read data", http.StatusInternalServerError)
				return
			}
			notes = append(notes, n)
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(notes)
	}
}


type UpdateNoteRequest struct {
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Tags        string  `json:"tags"` // Comma-separated
}

func UpdateNote(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		noteID := r.URL.Query().Get("id")
		if noteID == "" {
			http.Error(w, "Note ID is required", http.StatusBadRequest)
			return
		}

		var req UpdateNoteRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, "Invalid request body", http.StatusBadRequest)
			return
		}

		_, err := db.Exec(`
			UPDATE notes SET title = ?, content = ? WHERE id = ?
		`, req.Title, req.Description, noteID)

		if err != nil {
			http.Error(w, "Failed to update note", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]string{"message": "Note updated successfully"})
	}
}


func DeleteNote(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		noteID := r.URL.Query().Get("id")
		if noteID == "" {
			http.Error(w, "Note ID is required", http.StatusBadRequest)
			return
		}

		_, err := db.Exec(`DELETE FROM notes WHERE id = ?`, noteID)
		if err != nil {
			http.Error(w, "Failed to delete note", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]string{"message": "Note deleted successfully"})
	}
}

