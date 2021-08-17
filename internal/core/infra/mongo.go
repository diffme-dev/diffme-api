package infra

import (
	"context"
	"github.com/go-bongo/bongo"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

var (
	mongoClient *mongo.Client
)

var (
	Uri         = "mongodb://localhost:27017"
	mongoConfig = &bongo.Config{
		ConnectionString: "localhost",
		Database:         "diffme",
	}
)

func NewBongoConnection() (*bongo.Connection, error) {
	client, err := bongo.Connect(mongoConfig)

	if err != nil {
		return nil, err
	}

	return client, nil
}

func NewMongoConnection() (*mongo.Client, error) {
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(Uri+"/diffme"))

	if err != nil {
		log.Printf("Error connecting to mongo %s", err)
		return nil, err
	}

	return client, nil
}
