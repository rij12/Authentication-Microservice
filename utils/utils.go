package utils

import (
	"encoding/json"
	"errors"
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/rij12/Authentication-Microservice/models"
	"log"
	"net/http"
	"os"
)

func RespondWithErrorWithMessage(w http.ResponseWriter, status int, message string) {
	w.WriteHeader(status)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(message)
}

func RespondWithError(w http.ResponseWriter, status int) {
	w.WriteHeader(status)
}

func ResponseJSON(w http.ResponseWriter, status int, data interface{}) {
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(data)
}

// Checks if a user is valid
func ValidateUser(user models.User) error {

	if user.Password == "" || user.Email == "" {
		return errors.New("username or Password is empty")
	}
	return nil
}

// Generate JWT Token
func GenerateToken(user models.User) (string, error) {

	var err error
	secret := os.Getenv("JWT_SECRET")

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email": user.Email,
		"iss":   "course",
	})

	tokenString, err := token.SignedString([]byte(secret))

	if err != nil {
		log.Fatal(err)
	}

	return tokenString, nil

}
