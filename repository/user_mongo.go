package userRepository

import (
	"context"

	"github.com/rij12/Authentication-Microservice/models"
	"go.mongodb.org/mongo-driver/bson"
)

type UserRepository struct{}

func (u UserRepository) SaveUser(user models.User) (models.User, error) {

	_, err := db.Database("user").Collection("users").InsertOne(context.Background(), user)

	if err != nil {
		return models.User{}, err
	}

	return user, nil
}

func (u UserRepository) GetUserByEmail(user models.User) (models.User, error) {

	var userResult models.User
	err := uc.DatabaseClient.Database("user").Collection("users").FindOne(context.Background(), bson.M{"email": email[0]}).Decode(&user)

	if err != nil {
		return models.User{}, err
	}

	return userResult, nil
}
