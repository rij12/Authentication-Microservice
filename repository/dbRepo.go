package repository

import (
	"context"
	"errors"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

//var db *mongo.Client

type MongoRepository struct {
	mongoClient *mongo.Client
}

func (dbs *MongoRepository) Init(username string, password string, url string, port int) *MongoRepository {
	dbs.mongoClient = dbs.ConnectDB(username, password, url, port)
	return dbs
}

func (dbs *MongoRepository) ConnectDB(username string, password string, url string, port int) *mongo.Client {

	connectionString := fmt.Sprintf("mongodb://%s:%s@%s:%d/?authSource=admin", username, password, url, port)

	client, err := mongo.NewClient(options.Client().ApplyURI(connectionString))

	if err != nil {
		log.Fatal("Failed to create MongoRepository Client with error: ", err)
	}

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)

	if err != nil {
		log.Fatal("Failed to connect MongoRepository Client with error: ", err)
	}

	err = client.Ping(ctx, readpref.Primary())

	if err != nil {
		log.Fatal("Failed to ping MongoRepository Client with error: ", err)
	}
	return client
}

func (dbs *MongoRepository) PingDb() error {

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err := dbs.mongoClient.Connect(ctx)

	if err != nil {
		log.Fatal(err)
	}

	err = dbs.mongoClient.Ping(ctx, readpref.Primary())

	if err != nil {
		log.Fatal(err)
	}

	return err
}

func (dbs *MongoRepository) ListDatabases() []string {

	if dbs.mongoClient == nil {
		log.Fatal("MongoRepository Client is null")
		return []string{}
	}

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	databases, _ := dbs.mongoClient.ListDatabaseNames(ctx, bson.M{})

	fmt.Println(databases)

	return databases

}

func (dbs *MongoRepository) GetDb() (*mongo.Client, error) {
	if dbs.mongoClient == nil {
		return &mongo.Client{}, errors.New("mongo Client nil")
	}
	return dbs.mongoClient, nil
}

func (dbs *MongoRepository) CloseConnection() error {
	if dbs.mongoClient == nil {
		return errors.New("could not close mongo connection")
	}
	dbs.mongoClient.Disconnect(context.Background())
	return nil
}
