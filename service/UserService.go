package service

import (
	"errors"

	"github.com/rij12/models/models"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserService struct {
	DatabaseService *mongo.Client
}

func (userService UserService) SaveUser(user models.User) (models.User, error) {

	if user == nil {
		return nil, errors.New("User was nil")
	}

	userService.DatabaseService.Database("user").Collection("user").InsertOne()

}
