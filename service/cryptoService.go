package service

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/rij12/Authentication-Microservice/models"
	"github.com/rij12/Authentication-Microservice/utils"
	"log"
	"net/http"
	"strings"
	"time"
)

type CryptoService struct {
	config models.Config
}

func (cryptoService *CryptoService) Init(config models.Config) {
	cryptoService.config = config
}

// Generate JWT Token
func (cryptoService *CryptoService) GenerateToken(user models.User) (string, error) {

	var err error

	secret := cryptoService.config.Crypto.JWTSecret

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email": user.Email,
		"iss":   cryptoService.config.Crypto.JWTIssuer,
		"iat":   time.Now().Unix(),
		"exp":   time.Now().Add(time.Minute * cryptoService.config.Crypto.JWTExpire).Unix(),
	})

	tokenString, err := token.SignedString([]byte(secret))

	if err != nil {
		logger.Error("Error in generating JWT Token")
		log.Fatal(err)
	}

	logger.Debug(fmt.Sprintf("Token Service: Issued Token: %s at %s", token.Raw, time.Now()))

	return tokenString, nil
}

func (cryptoService *CryptoService) TokenVerifyMiddleWare(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		bearerToken := strings.Split(authHeader, " ")
		secret := cryptoService.config.Crypto.JWTSecret

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
				utils.RespondWithErrorWithMessage(w, http.StatusUnauthorized, errorMessage)
				return
			}
			if token.Valid {
				next.ServeHTTP(w, r)
			} else {
				utils.RespondWithError(w, http.StatusUnauthorized)
				return
			}
		} else {
			utils.RespondWithError(w, http.StatusUnauthorized)
			return
		}
	})
}
