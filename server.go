package main

import (
	"context"
	"fmt"
	"log"
	"reflect"


	// "go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func Server() *mongo.Client {
	clientOptions := options.Client().ApplyURI("mongodb://localhost:2000")

	// Connect to MongoDB
	Client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatal(err)
	}

	// Check the connection
	err = Client.Ping(context.TODO(), nil)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB!")
	fmt.Println("Connected to MongoDB!", reflect.TypeOf(Client))
	return Client
}