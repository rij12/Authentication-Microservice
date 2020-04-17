package service

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var db *mongo.Client

type DatabaseService struct {
	connectionString string
}

func (dbs *DatabaseService) init() *DatabaseService {
	return &DatabaseService{}
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

	db = client

	return client

}

func (dbs *DatabaseService) PingDb() error {

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err := db.Connect(ctx)

	if err != nil {
		log.Fatal(err)
	}

	err = db.Ping(ctx, readpref.Primary())

	if err != nil {
		log.Fatal(err)
	}

	return err
}

func (dbs *DatabaseService) ListDatabases() []string {

	if db == nil {
		log.Fatal("Database Client is null")
		return
	}

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	databases, _ := db.ListDatabaseNames(ctx, bson.M{})

	fmt.Println(databases)

	return databases

}

func (dbs *DatabaseService) getDb() *mongo.Client {
	return db
}
