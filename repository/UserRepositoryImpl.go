package repository

import (
	"context"
	"errors"
	"github.com/alexcesaro/log/stdlog"
	"github.com/rij12/Authentication-Microservice/models"
	"go.mongodb.org/mongo-driver/bson"
)

var logger = stdlog.GetFromFlags()

type UserRepositoryImpl struct{
	MongoRepository *MongoRepository
}

func (u UserRepositoryImpl) SaveUser(user models.User) (models.User, error) {

	repo, err := u.MongoRepository.GetDb()

	if err != nil {
		return models.User{}, errors.New(err.Error())
	}

	if u.checkUserExist(user) {
		logger.Info("UserRepositoryImpl:SaveUser: User already exists")
		return models.User{}, errors.New("user already exists")
	}

	_, err = repo.Database("user").Collection("users").InsertOne(context.Background(), user)

	if err != nil {
		logger.Error("UserRepositoryImpl:SaveUser: %s", err.Error())
		return models.User{}, errors.New("could not save user")
	}
	return user, nil
}

func (u UserRepositoryImpl) GetUserByEmail(email string) (models.User, error) {

	repo, err := u.MongoRepository.GetDb()

	if err != nil {
		return models.User{}, errors.New(err.Error())
	}

	var userResult models.User
	err = repo.Database("user").Collection("users").FindOne(context.Background(), bson.M{"email": email}).Decode(&userResult)

	if err != nil {
		logger.Error("UserRepositoryImpl:GetUserByEmail: %s", err.Error())
		return models.User{}, err
	}
	return userResult, nil
}

func (u UserRepositoryImpl) GetUser(user models.User) (models.User, error) {

	repo, err := u.MongoRepository.GetDb()

	if err != nil {
		logger.Error("UserRepositoryImpl:GetUser: %s", err.Error())
		return models.User{}, errors.New(err.Error())
	}

	var userResult models.User
	err = repo.Database("user").Collection("users").FindOne(context.Background(), user).Decode(&userResult)

	if err != nil {
		logger.Error("UserRepositoryImpl:GetUser: %s", err.Error())
		return models.User{}, err
	}
	return userResult, nil
}

func (u UserRepositoryImpl) checkUserExist(user models.User) bool {
	userResult, _ := u.GetUserByEmail(user.Email)
	if userResult.Email == user.Email {
		return true
	}
	return false
}
