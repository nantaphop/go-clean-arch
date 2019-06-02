package core

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// GetMongo return a mongodb client
func GetMongo() (*mongo.Client, error) {
	clientOptions := options.Client().ApplyURI("mongodb://root:password@mongo:27017")
	ctx, cancel := context.WithTimeout(context.Background(), 10*1000)
	defer cancel()
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	// Check the connection
	err = client.Ping(context.TODO(), nil)

	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	return client, nil
}
