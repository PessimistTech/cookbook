package db

import (
	"context"
	"fmt"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var client *mongo.Client

func InitDB() (*mongo.Client, error) {
	url := os.Getenv("MONGODB_URL")
	cl, err := mongo.Connect(context.Background(), options.Client().ApplyURI(url))
	if err != nil {
		return &mongo.Client{}, fmt.Errorf("Failed to connect to db: %+v", err)
	}

	client = cl

	return cl, nil

}

func GetClient() *mongo.Client {
	return client
}

func GetCollection(client *mongo.Client, database, collection string) *mongo.Collection {
	cltn := client.Database(database).Collection(collection)
	return cltn
}
