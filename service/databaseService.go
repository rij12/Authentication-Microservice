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

var databasesHandle *mongo.Client

type DatabaseService struct {
	connectionString string
	client           *mongo.Client
}

func (dbs *DatabaseService) newDatabaseService() *DatabaseService {
	return &DatabaseService{}
}

func (dbs *DatabaseService) createUser(user models.User) {

}

func (dbs *DatabaseService) updateUser(user models.User) {

}

func (dbs *DatabaseService) deleteUser(user models.User) {

}

func (dbs *DatabaseService) getUserByEmail(email string) {

}

func (dbs *DatabaseService) getUserByID(id string) {

}

func (dbs *DatabaseService) ConnectDB(username string, password string, url string, port int) *mongo.Client {

	dbs.connectionString = fmt.Sprintf("mongodb://%s:%s@%s:%d/?authSource=admin", username, password, url, port)

	client, err := mongo.NewClient(options.Client().ApplyURI(dbs.connectionString))

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

	dbs.client = client
	databasesHandle = client
	return client

}

func (dbs *DatabaseService) PingDb() error {

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

func (dbs *DatabaseService) ListDatabases() {

	if dbs.client == nil {
		log.Fatal("Database Client is null")
		return
	}

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	databases, _ := dbs.client.ListDatabaseNames(ctx, bson.M{})

	fmt.Println(databases)

}
