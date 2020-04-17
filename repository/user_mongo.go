package repository

import (
	"context"

	"github.com/alexcesaro/log/stdlog"
	"github.com/rij12/Authentication-Microservice/models"
	"go.mongodb.org/mongo-driver/bson"
)

var logger = stdlog.GetFromFlags()

type UserRepository struct{}

func (u UserRepository) SaveUser(user models.User) (models.User, error) {

	database := Database{}
	repo := database.getDb()

	_, err := repo.Database("user").Collection("users").InsertOne(context.Background(), user)

	if err != nil {
		logger.Warning("UserRepository: Could not save user: %s to database", user)
		return models.User{}, err
	}
	return user, nil
}

func (u UserRepository) GetUserByEmail(email string) (models.User, error) {

	database := Database{}
	repo := database.getDb()

	var userResult models.User
	err := repo.Database("user").Collection("users").FindOne(context.Background(), bson.M{"email": email}).Decode(&userResult)

	if err != nil {
		logger.Warning("UserRepository: Could not find user with email: %s in database", email)
		return models.User{}, err
	}
	return userResult, nil
}

func (u UserRepository) GetUser(user models.User) (models.User, error) {

	database := Database{}
	repo := database.getDb()

	var userResult models.User
	err := repo.Database("user").Collection("users").FindOne(context.Background(), user).Decode(&userResult)

	if err != nil {
		logger.Warning("UserRepository: Could not get user: %s from database", user)
		return models.User{}, err
	}
	return userResult, nil
}
