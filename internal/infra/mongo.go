package infra

import (
	"github.com/go-bongo/bongo"
	"go.mongodb.org/mongo-driver/mongo"
)

var (
	mongoClient *mongo.Client
)

var (
	mongoConfig = &bongo.Config{
		ConnectionString: "localhost",
		Database:         "diffme",
	}
)

func NewMongoConnection() (*bongo.Connection, error) {

	// Replace the uri string with your MongoDB deploym

	client, err := bongo.Connect(mongoConfig)

	if err != nil {
		return nil, err
	}

	return client, nil
}
