package utils

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/alexcesaro/log/stdlog"
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/rij12/Authentication-Microservice/models"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
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

// Generate JWT Token
func GenerateToken(user models.User) (string, error) {

	var err error
	secret := os.Getenv("JWT_KEY")

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email": user.Email,
		"iss":   "RPJ_AUTH",
		"iat":   time.Now().Unix(),
		"exp":   time.Now().Add(time.Minute * time.Duration(1)).Unix(),
	})

	tokenString, err := token.SignedString([]byte(secret))

	if err != nil {
		logger.Error("Error in generating JWT Token")
		log.Fatal(err)
	}

	logger.Debug(fmt.Sprintf("Token Service: Issued Token: %s at %s", token.Raw, time.Now()))

	return tokenString, nil
}

func TokenVerifyMiddleWare(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		bearerToken := strings.Split(authHeader, " ")
		secret := os.Getenv("JWT_KEY")

		if len(bearerToken) == 2 {
			authToken := bearerToken[1]

			token, err := jwt.Parse(authToken, func(token *jwt.Token) (interface{}, error) {
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					logger.Error("Token Service: Error Verifying JWT Algorithm")
					return nil, fmt.Errorf("internal Server Error")
				}
				return []byte(secret), nil
			})
			if err != nil {
				errorMessage := err.Error()
				logger.Info("Token Service: ", errorMessage)
				RespondWithErrorWithMessage(w, http.StatusUnauthorized, errorMessage)
				return
			}
			if token.Valid {
				next.ServeHTTP(w, r)
			} else {
				RespondWithError(w, http.StatusUnauthorized)
				return
			}
		} else {
			RespondWithError(w, http.StatusUnauthorized)
			return
		}
	})
}
