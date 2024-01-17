package common

import (
	"context"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var db *mongo.Database

func GetDBCollection(col string) *mongo.Collection {
	return db.Collection(col)
}

func InitDB() error {
	uri := os.Getenv("MONGODB_URI")

	if uri == "" {
		uri = "mongodb://localhost:27017/gomongodb"
	}


	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(uri))
	if err != nil {
		return err
	}

	db = client.Database("go_demo")

	return nil
}

func CloseDB() error {
	return db.Client().Disconnect(context.Background())
}