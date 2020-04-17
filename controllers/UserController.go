package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/rij12/Authentication-Microservice/utils"
	"net/http"

	"github.com/alexcesaro/log/stdlog"
	"github.com/google/uuid"
	"github.com/rij12/Authentication-Microservice/models"
	"github.com/rij12/Authentication-Microservice/service"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

var logger = stdlog.GetFromFlags()

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

	decoder := json.NewDecoder(r.Body)
	var user models.User
	err := decoder.Decode(&user)

	if err != nil {
		logger.Warning("LoginController: Could decode JSON")
		utils.RespondWithError(w, http.StatusInternalServerError)
		return
	}

	// User Service
	userService := service.UserService{}
	jwt, serviceError := userService.Login(user)

	if serviceError != nil {
		utils.RespondWithError(w, http.StatusUnauthorized)
		return
	}

	// Set return response
	utils.ResponseJSON(w, http.StatusOK, jwt)
}

func (uc *UserController) RegisterController(w http.ResponseWriter, r *http.Request) {

	//TODO
	// Check if user already is in database

	// Parse body into a User
	decoder := json.NewDecoder(r.Body)
	var user models.User
	err := decoder.Decode(&user)
	user.UserID = uuid.New().String()

	fmt.Println(user)

	validationError := utils.ValidateUser(user)
	if validationError != nil {
		logger.Warning("RegisterController: User missing fields %s", user)
		utils.RespondWithErrorWithMessage(w, http.StatusBadRequest, validationError.Error())
		return
	}

	if err != nil {
		logger.Warning("RegisterController: Could decode User %s", user)
		utils.RespondWithErrorWithMessage(w, http.StatusBadRequest, err.Error())
		return
	}

	// Generate Crypto
	hash, CryptoError := bcrypt.GenerateFromPassword([]byte(user.Password), 10)
	user.Password = string(hash)
	if CryptoError != nil {
		utils.RespondWithError(w, http.StatusBadRequest)
		return
	}

	// User Service
	userService := service.UserService{}
	_, userServiceError := userService.RegisterUser(user)
	if userServiceError != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	utils.ResponseJSON(w, http.StatusCreated, nil)
}

func (uc *UserController) ProtectedEndpointTest(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Protected Endpoint")
}

func (uc *UserController) GetUserByEmailController(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	email, err := r.URL.Query()["email"]

	if !err || len(email[0]) < 1 {
		logger.Warning("GetUserByEmailController: Could not find URL Param 'email'")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	userService := service.UserService{}
	user, userServiceError := userService.GetUserByEmail(email[0])

	if userServiceError != nil {
		utils.RespondWithErrorWithMessage(w, http.StatusNotFound, userServiceError.Error())
		return
	}

	utils.ResponseJSON(w, http.StatusOK, user)
}
