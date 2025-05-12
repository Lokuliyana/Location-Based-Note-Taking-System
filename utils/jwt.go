package utils

import (
	"errors"
	"net/http"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
)
var jwtSecret = []byte("9f2d1a21a7b4e01c872d234d3ffba18d9ae4d1a5a3f1c59e4c1470f97f4fd22a") 

// Load the JWT key from an environment variable
var jwtKey = []byte(os.Getenv("JWT_SECRET"))

// GenerateJWT generates a signed JWT token for a user
func GenerateJWT(userID uint) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": userID,
		"exp":     time.Now().Add(24 * time.Hour).Unix(),
	})

	return token.SignedString(jwtKey)
}

// ParseJWT validates and parses the JWT token to extract claims
func ParseJWT(tokenString string) (*jwt.Token, jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})

	if err != nil || !token.Valid {
		return nil, nil, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, nil, err
	}

	return token, claims, nil
}

// Example handler (for documentation/testing purposes)
func SomeProtectedHandler(w http.ResponseWriter, r *http.Request) {
	tokenString := r.Header.Get("Authorization")
	if tokenString == "" {
		http.Error(w, "Missing token", http.StatusUnauthorized)
		return
	}

	// Directly call ParseJWT (no utils.ParseJWT inside utils package)
	_, _, err := ParseJWT(tokenString)
	if err != nil {
		http.Error(w, "Invalid token", http.StatusUnauthorized)
		return
	}

	// Get the user ID from claims, if needed
	// For demonstration, we're not using it to avoid "unused variable" error
	// userID := claims["user_id"]

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Token is valid"))
}

// VerifyJWT verifies and parses a JWT token string
func VerifyJWT(tokenString string) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Validate the algorithm
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return jwtSecret, nil
	})
	return token, err
}
