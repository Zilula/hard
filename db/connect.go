package db

import (
	"context"
	"fmt"
	"log"
	// "encoding/json"
	// "net/http"

	// "go.mongodb.org/mongo-driver/bson"
	// "github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)


func CreateConnection() *mongo.Database {

	// Rest of the code will go here
	// Set client options
	clientOptions := options.Client().ApplyURI("mongodb://localhost:2000")

	// Connect to MongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatal(err)
	}

	// Check the connection
	err = client.Ping(context.TODO(), nil)

	if err != nil {
		log.Fatal(err)
	}
	db := client.Database("Pokemon")

	fmt.Println("Connected to MongoDB!")
	// creates the collection
	collection := db.Collection("trainers")

	// create a new trainer
	ash := Trainer{"Ash", 10, "Pallet Town"}
	misty := Trainer{"Misty", 10, "Pallet Town"}

	// insert the two trainers into the collection
	collection.InsertOne(context.TODO(), misty)
	collection.InsertOne(context.TODO(), ash)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Two trainers added to the collection")

	return db
}
