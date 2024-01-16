package config

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var client *mongo.Client

func init() {
    client = ConnectDB()
}

func ConnectDB() *mongo.Client {
	cfg := GetConfig()

	if cfg.MONGODB_URI == "" {
		log.Fatal("MONGODB_URI is not set")
	}

	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(cfg.MONGODB_URI))
	if err != nil {
		panic(err)
	}

	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to Mongodb")

	return client
}

func GetCollection(collectionName string) *mongo.Collection {
	collection := client.Database("gomongodb").Collection(collectionName)
	return collection
}