package repository

import (
	"context"

	"github.com/rij12/Authentication-Microservice/models"
	"go.mongodb.org/mongo-driver/bson"
)

type UserRepository struct{}

func (u UserRepository) SaveUser(user models.User) (models.User, error) {

	database := Database{}
	repo := database.getDb()

	_, err := repo.Database("user").Collection("users").InsertOne(context.Background(), user)

	if err != nil {
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
		return models.User{}, err
	}

	return userResult, nil
}
