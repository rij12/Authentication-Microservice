package controllers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/google/uuid"
	"github.com/rij12/Authentication-Microservice/models"
	"github.com/rij12/Authentication-Microservice/service"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

type UserController struct {
	DatabaseClient *mongo.Client
	UserService    *service.UserService
}

func (uc *UserController) LoginController(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Login Controller")
}

func (uc *UserController) RegisterController(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var user models.User
	err := decoder.Decode(&user)
	user.UserID = uuid.New().String()
	if err != nil {
		log.Fatal(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Hash the password
	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), 10)

	if err != nil {
		log.Fatal(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	user.Password = string(hash)

	// User Service
	result := service.UserService.RegisterUser(user)

}

func (uc *UserController) ProtectedEndpointTest(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Protected Endpoint")
}

func (uc *UserController) GetUserByEmailController(w http.ResponseWriter, r *http.Request) {

	fmt.Println("User by email handeler hit!")

	w.Header().Set("Content-Type", "application/json")

	email, err := r.URL.Query()["email"]

	if !err || len(email[0]) < 1 {
		log.Warning("Url Param 'email' is missing")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	user, err := service.UserService.GetUserByEmail(email)
	json.NewEncoder(w).Encode(user)
}

func (uc *UserController) GetDbHealth(w http.ResponseWriter, r *http.Request) {
	return w.WriteHeader(http.StatusOK)
}

func handleError(err error) http.Response {

	// TODO!
}
