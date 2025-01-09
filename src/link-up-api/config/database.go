package config

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var DB *mongo.Client

func ConnectDatabase(uri string) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	clientOptions := options.Client().ApplyURI(uri).SetAuth(options.Credential{
		Username: "root",
		Password: "root",
	})

	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatalf("Failed to connect to MongoDB: %v", err)
	}

	DB = client
	log.Println("Connected to MongoDB!")
}

func GetCollection(database, collection string) *mongo.Collection {
	return DB.Database(database).Collection(collection)
}
