package service

import (
	"errors"
	"fmt"
	"github.com/alexcesaro/log/stdlog"
	"github.com/rij12/Authentication-Microservice/models"
	"github.com/rij12/Authentication-Microservice/repository"
	"github.com/rij12/Authentication-Microservice/utils"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
	"log"
)

var logger = stdlog.GetFromFlags()

type UserService struct {
	DatabaseService *mongo.Client
}

func (userService UserService) RegisterUser(user models.User) (models.User, error) {

	userRepo := new(repository.UserRepository)
	user, err := userRepo.SaveUser(user)

	if err != nil {
		logger.Warning("UserService: Could not register user: %s", user)
		return models.User{}, err
	}

	return user, nil
}

func (userService UserService) GetUserByEmail(email string) (models.UserResult, error) {

	userRepo := repository.UserRepository{}
	user, err := userRepo.GetUserByEmail(email)

	if err != nil {
		logger.Warning("UserService: Could not get user with email: %s", email)
		return models.UserResult{}, errors.New(fmt.Sprintf("Can not find user with email: %s", email))
	}
	stripedUser := models.UserResult{Email: user.Email, UserID: user.UserID}
	return stripedUser, nil
}

func (userService UserService) Login(user models.User) (models.JWT, error) {

	userRepo := repository.UserRepository{}
	userFromDatabase, err := userRepo.GetUserByEmail(user.Email)

	if err != nil {
		logger.Warning("UserService: Could not get user %s from database", user)
		return models.JWT{}, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(userFromDatabase.Password), []byte(user.Password))

	if err != nil {
		logger.Warning("UserService: Passwords do not match")
		return models.JWT{}, err
	}

	// Generate JWT Token
	token, err := utils.GenerateToken(user)

	if err != nil {
		logger.Critical("UserService: Can not generate JWT Token")
		log.Fatal(err)
	}

	jwt := models.JWT{Token: token}

	return jwt, nil

}
