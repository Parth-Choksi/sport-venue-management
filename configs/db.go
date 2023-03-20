package configs

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func Connect() *mongo.Client {
	// Database Config
	// mongodb://127.0.0.1:27017
	// mongodb://127.0.0.1:27017/sports-venue-management?readPreference=primary&appname=MongoDB%20Compass&directConnection=true&ssl=false?retryWrites=true&w=majority
	clientOptions := options.Client().ApplyURI("mongodb://127.0.0.1:27017")
	client, err := mongo.NewClient(clientOptions)

	// setup to context required by mongo

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)

	if err != nil {
		log.Fatal(err)
	}

	defer cancel()

	// ping your db connection used to verify that mongoDB server is available and responsive.
	err = client.Ping(context.Background(), readpref.Primary())

	if err != nil {
		log.Fatal("could not connect to the database", err)
	} else {
		log.Println("Connected")
	}

	return client
}

var DB *mongo.Client = Connect()

func GetCollection(client *mongo.Client, collectionName string) *mongo.Collection {
	collection := client.Database("sports-venue-management").Collection(collectionName)
	return collection
}
