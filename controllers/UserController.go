package controllers

import (
	"fmt"
	"net/http"

	"github.com/rij12/Authentication-Microservice/models"
)

type UserController struct{}

func (uc UserController) LoginController(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Login Controller")
}

func (uc UserController) RegisterController(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Register Controller")
}

func (uc UserController) ProtectedEndpointTest(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Protected Endpoint")
}

func (uc UserController) GetUserController(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get user Controller endPoint hit")
}

func (uc UserController) GetDbHealth(w http.ResponseWriter, r *http.Request) {
	fmt.Println("GetDbHealth Called.")

	err := models.Database.PingDb()

	if err != nil {
		http.Error(w, "Can not connect to MongoDB", http.StatusInternalServerError)
	}

}
