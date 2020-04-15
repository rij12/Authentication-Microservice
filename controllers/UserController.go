package controllers

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/google/uuid"
	"github.com/rij12/Authentication-Microservice/models"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserController struct {
	DatabaseClient *mongo.Client
}

func (uc *UserController) LoginController(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Login Controller")
}

func (uc *UserController) RegisterController(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var user models.User
	err := decoder.Decode(&user)
	user.ID = uuid.New().String()
	if err != nil {
		panic(err)
	}
	log.Println(user)
	uc.DatabaseClient.Database("user").Collection("users").InsertOne(context.TODO(), user)
}

func (uc *UserController) ProtectedEndpointTest(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Protected Endpoint")
}

func (uc *UserController) GetUserController(w http.ResponseWriter, r *http.Request) {

}

func (uc *UserController) GetDbHealth(w http.ResponseWriter, r *http.Request) {
	fmt.Println("GetDbHealth Called.")

}
