package repository

import "github.com/rij12/Authentication-Microservice/models"

type UserRepository interface {

	SaveUser(user models.User) (models.User, error)
	GetUserByEmail(email string) (models.User, error)
	GetUser(user models.User) (models.User, error)
	checkUserExist(user models.User) bool
}


