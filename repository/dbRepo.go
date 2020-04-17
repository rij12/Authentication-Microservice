package repository

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

type Database struct {
	connectionString string
}

func (dbs *Database) init() *Database {
	return &Database{}
}

func (dbs *Database) ConnectDB(username string, password string, url string, port int) *mongo.Client {

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

func (dbs *Database) PingDb() error {

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

func (dbs *Database) ListDatabases() []string {

	if db == nil {
		log.Fatal("Database Client is null")
		return []string{}
	}

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	databases, _ := db.ListDatabaseNames(ctx, bson.M{})

	fmt.Println(databases)

	return databases

}

func (dbs *Database) getDb() *mongo.Client {
	return db
}
