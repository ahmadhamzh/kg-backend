package database

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ConnectToDB() (*mongo.Collection, error) {
	// Set up MongoDB client
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb+srv://kgProject:a2dQcm9qZWN0Cg==@kgdatabase.sc7u6rk.mongodb.net/?retryWrites=true&w=majority"))
	if err != nil {
		log.Fatal(err)
	}

	// Connect to MongoDB
	err = client.Connect(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	// Get handle to database and collection
	collection := client.Database("kgDataBase").Collection("teachers")

	return collection, nil
}
