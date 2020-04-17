package service

import (
	"fmt"
	"github.com/rij12/Authentication-Microservice/models"
	"github.com/rij12/Authentication-Microservice/repository"

	"go.mongodb.org/mongo-driver/mongo"
)

type UserService struct {
	DatabaseService *mongo.Client
}

func (userService UserService) RegisterUser(user models.User) (models.User, error) {

	userRepo := new(repository.UserRepository)
	user, err := userRepo.SaveUser(user)

	if err != nil {
		return models.User{}, err
	}

	return user, nil
}

func (userService UserService) GetUserByEmail(email string) (models.User, error) {

	fmt.Println(email)

	userRepo := repository.UserRepository{}
	user, err := userRepo.GetUserByEmail(email)

	fmt.Println(user)

	if err != nil {
		return models.User{}, err
	}
	return user, nil
}

func (userService UserService) Login(email string, hashedPassword string) (models.JWT, error) {
	return models.JWT{}, nil
}
