package service

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/rij12/Authentication-Microservice/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type DatabaseService struct {
	connectionString string
	client           *mongo.Client
	ctx              *context.Context
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

func (dbs DatabaseService) ConnectDB(username string, password string, url string, port int) *mongo.Client {

	connectionString := fmt.Sprintf("mongodb://%s:%s@%s:%d/?authSource=admin", username, password, url, port)
	dbs.connectionString = connectionString

	client, err := mongo.NewClient(options.Client().ApplyURI(connectionString))

	if err != nil {
		log.Fatal(err)
	}

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	dbs.ctx = &ctx
	err = client.Connect(ctx)

	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(ctx, readpref.Primary())

	if err != nil {
		log.Fatal(err)
	}

	dbs.client = client
	return client

}

func (dbs DatabaseService) PingDb() error {

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err := dbs.client.Connect(ctx)

	if err != nil {
		log.Fatal(err)
	}

	err = dbs.client.Ping(ctx, readpref.Primary())

	if err != nil {
		log.Fatal(err)
	}

	return err

}

func (dbs DatabaseService) ListDatabases() {

	databases, _ := dbs.client.ListDatabaseNames(*dbs.ctx, bson.M{})

	fmt.Println(databases)

}
