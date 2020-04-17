package service

import (
	"github.com/rij12/Authentication-Microservice/models"
	"github.com/rij12/Authentication-Microservice/repository"

	"go.mongodb.org/mongo-driver/mongo"
)

type UserService struct {
	DatabaseService *mongo.Client
}

func (userService UserService) RegisterUser(user models.User) (models.User, error) {

	userRepo := new(repository.UserRepository)
	userRepo.Register(user)
	return user, nil
}

func (userService UserService) GetUserByEmail(email string) (models.User, models.Error) {

	userRepo := new(repository.UserRepository)
	userRepo.FindUserByEmail(user)
	return user, nil
}

func (userService UserService) Login(email string, hashedPassword string) (models.JWT, error) {

}
