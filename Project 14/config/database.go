package config

import (
    "context"
    "time"
	"os"

    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
)

type DatabaseCollections struct {
    Dogs *mongo.Collection
}

var Collections DatabaseCollections
var Client *mongo.Client

func ConnectDatabase() error {

	uriDB := os.Getenv("MONGODB_URI")
	if uriDB == "" {
		uriDB = "mongodb://localhost:27017/"
	}

    client, err := mongo.NewClient(options.Client().ApplyURI(uriDB))
    ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
    defer cancel()

    err = client.Connect(ctx)
    if err != nil {
        return err
    }

    db := client.Database("DEV")
    dogsCollection := db.Collection("dogs")
	
    Collections = DatabaseCollections{
        Dogs: dogsCollection,
    }
    Client = client
    return nil
}