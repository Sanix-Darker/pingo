package main

import (
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// MongoClient
func MongoClient() *mongo.Client {
	var mongoClient, err = mongo.Connect(nil, options.Client().ApplyURI(MONGO_HOST))
	if err != nil {
		log.Fatal("Error: creating mongo-client")
	}

	return mongoClient
}
