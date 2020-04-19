package utils

import (
	"encoding/json"
	"errors"
	"github.com/alexcesaro/log/stdlog"
	"github.com/rij12/Authentication-Microservice/models"
	"net/http"
)

var logger = stdlog.GetFromFlags()

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

func Response(w http.ResponseWriter, status int) {
	w.WriteHeader(status)
}

// Checks if a user is valid
func ValidateUser(user models.User) error {

	if user.Password == "" || user.Email == "" {
		return errors.New("username or Password is empty")
	}
	return nil
}
