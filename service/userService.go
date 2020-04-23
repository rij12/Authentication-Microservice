package service

import (
	"errors"
	"fmt"
	"github.com/alexcesaro/log/stdlog"
	"github.com/rij12/Authentication-Microservice/models"
	"github.com/rij12/Authentication-Microservice/repository"
	"golang.org/x/crypto/bcrypt"
	"log"
)

var logger = stdlog.GetFromFlags()

type UserService struct {
	ConfigurationService ConfigurationService
	CryptoService CryptoService
	UserRepository repository.UserRepository
}

func (userService UserService) RegisterUser(user models.User) (models.User, error) {

	logger.Info(fmt.Sprintf("UserService:RegisterUser: Attempting to register user with email: %s", user.Email))

	user, err := userService.UserRepository.SaveUser(user)

	if err != nil {
		logger.Error(fmt.Sprintf("UserService:RegisterUser: %s", err.Error()))
		return models.User{}, err
	}

	return user, nil
}

func (userService UserService) GetUserByEmail(email string) (models.UserResult, error) {

	logger.Info(fmt.Sprintf("UserService:GetUserByEmail: Attempting to get user with email: %s", email))

	user, err := userService.UserRepository.GetUserByEmail(email)

	if err != nil {
		logger.Error(fmt.Sprintf("UserService: Could not get user with email: %s", email))
		return models.UserResult{}, errors.New(fmt.Sprintf("Can not find user with email: %s", email))
	}
	stripedUser := models.UserResult{Email: user.Email, UserID: user.UserID}
	return stripedUser, nil
}

func (userService UserService) Login(user models.User) (models.JWT, error) {

	logger.Info("UserService:Login: Attempting to login user with ID:", user.UserID)

	userFromDatabase, err := userService.UserRepository.GetUserByEmail(user.Email)

	if err != nil {
		logger.Error(fmt.Sprintf("UserService:Login: Could not get user %s from database", user))
		return models.JWT{}, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(userFromDatabase.Password), []byte(user.Password))

	if err != nil {
		logger.Error("UserService:Login: Passwords do not match")
		return models.JWT{}, err
	}

	// Generate JWT Token
	token, err := userService.CryptoService.GenerateToken(user)

	if err != nil {
		logger.Error("UserService:Login: Can not generate JWT Token")
		log.Fatal(err)
	}

	jwt := models.JWT{Token: token}

	return jwt, nil
}

