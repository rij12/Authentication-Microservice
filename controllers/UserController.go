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

// Login godoc
// @Summary Login
// @Description login by email and password
// @ID login
// @Accept  json
// @Produce  json
// @Param id path int true "Account ID"
// @Success 200 {object} model.Account
// @Header 200 {string} Token "qwerty"
// @Failure 400 {object} httputil.HTTPError
// @Failure 404 {object} httputil.HTTPError
// @Failure 500 {object} httputil.HTTPError
// @Router /accounts/{id} [get]
func (uc *UserController) LoginController(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Login Controller")
}

func (uc *UserController) RegisterController(w http.ResponseWriter, r *http.Request) {

	// Parse body into a User
	decoder := json.NewDecoder(r.Body)
	var user models.User
	err := decoder.Decode(&user)
	user.UserID = uuid.New().String()
	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), 10)
	user.Password = string(hash)

	if err != nil {
		log.Fatal(err)
		w.WriteHeader(http.StatusBadRequest)
	}

	if err != nil {
		log.Fatal(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// User Service
	userService := service.UserService{}
	_, err = userService.RegisterUser(user)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}

	w.WriteHeader(http.StatusCreated)

}

func (uc *UserController) ProtectedEndpointTest(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Protected Endpoint")
}

func (uc *UserController) GetUserByEmailController(w http.ResponseWriter, r *http.Request) {

	fmt.Println("User by email handeler hit!")

	w.Header().Set("Content-Type", "application/json")

	email, err := r.URL.Query()["email"]

	if !err || len(email[0]) < 1 {
		//log.("Url Param 'email' is missing")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	userService := service.UserService{}
	user, _ := userService.GetUserByEmail(email[0])
	json.NewEncoder(w).Encode(user)
}

func (uc *UserController) GetDbHealth(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

//func handleError(err error) http.Response {
//
//	// TODO!
//}
