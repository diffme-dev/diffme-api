package loaders

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"time"
)

var (
	mongoClient *mongo.Client
)

func ConnectMongo() *mongo.Client {
	println("starting mongo")

	// Replace the uri string with your MongoDB deployment's connection string.
	uri := "mongodb://localhost:27017/diffme?w=majority"

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	defer cancel()

	mongoClient, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))

	if err != nil {
		panic(err)
	}

	defer func() {
		println("disconnecting...")
		if err = mongoClient.Disconnect(ctx); err != nil {
			panic(err)
		}
	}()

	// Ping the primary
	if err := mongoClient.Ping(ctx, readpref.Primary()); err != nil {
		panic(err)
	}

	fmt.Println("Successfully connected and pinged.")

	return mongoClient
}
