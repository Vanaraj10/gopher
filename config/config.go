package config

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)
// Package config provides functionality to configure and connect to the MongoDB database.
// It exposes methods to establish a connection and retrieve collections from the database.
//
// Variables:
//   DB - A global variable holding the MongoDB client instance.
//
// Functions:
//   ConnectDB() - Establishes a connection to MongoDB using environment variables "MONGO_URI" and sets the global DB client.
//   GetCollection(name string) *mongo.Collection - Returns a collection from the connected MongoDB database specified by the "DB_NAME" environment variable.
var DB *mongo.Client

func ConnectDB() {
	mongoURI := os.Getenv("MONGO_URI")

	client, err := mongo.NewClient(options.Client().ApplyURI(mongoURI))
	
	if err != nil {
		log.Fatalf("Failed to create MongoDB client: %v", err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10 * time.Second)
	defer cancel()

	err = client.Connect(ctx)
	if err != nil {
		log.Fatalf("Failed to connect to MongoDB: %v", err)
	}
	DB = client
	fmt.Println("Connected to MongoDB successfully")
}

func GetCollection(name string) *mongo.Collection {
	dbName := os.Getenv("DB_NAME")
	return DB.Database(dbName).Collection(name)
}

