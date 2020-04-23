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
	"golang.org/x/crypto/bcrypt"
)

var logger = stdlog.GetFromFlags()

type UserController struct {
	UserService    *service.UserService
}


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
	jwt, serviceError := uc.UserService.Login(user)

	if serviceError != nil {
		logger.Warning(serviceError.Error())
		utils.RespondWithError(w, http.StatusUnauthorized)
		return
	}

	// Set return response
	utils.ResponseJSON(w, http.StatusOK, jwt)
}

func (uc *UserController) RegisterController(w http.ResponseWriter, r *http.Request) {

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
	_, userServiceError := uc.UserService.RegisterUser(user)
	if userServiceError != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	utils.Response(w, http.StatusCreated)
}

func (uc *UserController) ProtectedEndpointTest(w http.ResponseWriter, r *http.Request) {

	m := make(map[string]string)
	m["result"] = "super secret information"

	utils.ResponseJSON(w, http.StatusOK, m)
}

func (uc *UserController) GetUserByEmailController(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	email, err := r.URL.Query()["email"]

	if !err || len(email[0]) < 1 {
		utils.RespondWithError(w, http.StatusBadRequest)
		return
	}

	user, userServiceError := uc.UserService.GetUserByEmail(email[0])

	if userServiceError != nil {
		utils.RespondWithError(w, http.StatusNotFound)
		return
	}

	utils.ResponseJSON(w, http.StatusOK, user)
}
