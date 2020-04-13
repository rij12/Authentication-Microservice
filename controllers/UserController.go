package controllers

import (
	"fmt"
	"net/http"
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
