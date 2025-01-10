package db

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Global variables for client and database
var client *mongo.Client
var database *mongo.Database

// Connect to MongoDB and return the client
func ConnectMongoDB() (*mongo.Client, *mongo.Database, error) {
	uri := "mongodb+srv://test:test@cluster0.5fn5s.mongodb.net/"

	clientOptions := options.Client().ApplyURI(uri)
	var err error
	client, err = mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		return nil, nil, err
	}

	// Check the connection
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		return nil, nil, err
	}

	// Set the database
	database = client.Database("crud") // Replace with your database name
	fmt.Println("Connected to MongoDB!")

	return client, database, nil
}

// GetMongoDB returns the MongoDB client and database instance
func GetMongoDB() (*mongo.Client, *mongo.Database, error) {
	return ConnectMongoDB()
}
