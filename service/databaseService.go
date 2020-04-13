package service

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/rij12/Authentication-Microservice/models"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type DatabaseService struct {
	// db *database
}

func (dbs DatabaseService) createUser(user models.User) {

}

func (dbs DatabaseService) updateUser(user models.User) {

}

func (dbs DatabaseService) deleteUser(user models.User) {

}

func (dbs DatabaseService) getUserByEmail(email string) {

}

func (dbs DatabaseService) getUserByID(id string) {

}

func (dbs DatabaseService) ConnectDB(username string, password string, url string, port int) {

	connectionString := fmt.Sprintf("mongodb://%s:%s@%s:%d/?authSource=admin", username, password, url, port)

	client, err := mongo.NewClient(options.Client().ApplyURI(connectionString))

	if err != nil {
		log.Fatal(err)
	}

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	err = client.Connect(ctx)

	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(ctx, readpref.Primary())

	if err != nil {
		log.Fatal(err)
	}

}

func (dbs DatabaseService) pingDb(){
	err	= dbs.ConnectDB("username", ,"password", "url", "port")
}
