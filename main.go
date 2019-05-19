package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
	"github.com/gorilla/mux"
	"github.com/mongodb/mongo-go-driver/bson"
	"github.com/mongodb/mongo-go-driver/bson/primitive"
	"github.com/mongodb/mongo-go-driver/mongo"
)

var client *mongo.Client

type Person struct {
	ID        primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Firstname string             `json:"firstname,omitempty" bson:"firstname,omitempty"`
	Lastname  string             `json:"lastname,omitempty" bson:"lastname,omitempty"`
}

func CreatePersonEndpoint(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("content-type", "application/json")
	var person Person
	_ = json.NewDecoder(request.Body).Decode(&person)
	collection := client.Database("thepolyglotdeveloper").Collection("people")
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	result, _ := collection.InsertOne(ctx, person)
	json.NewEncoder(response).Encode(result)
}
func main() {
	fmt.Println("Starting the application...")
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	client, _ = mongo.Connect(ctx, "mongodb://localhost:2000")
	router := mux.NewRouter()
	router.HandleFunc("/person", CreatePersonEndpoint).Methods("POST")
	router.HandleFunc("/people", GetPeopleEndpoint).Methods("GET")
	router.HandleFunc("/person/{id}", GetPersonEndpoint).Methods("GET")
	log.Fatal(http.ListenAndServe(":12345", router))
}


////////////////////////////////////////////////////////////////////////////////////
// package main

// import (
// 	"context"
// 	"fmt"
// 	"log"
// 	"encoding/json"
// 	"net/http"

// 	// "go.mongodb.org/mongo-driver/bson"
// 	"github.com/gorilla/mux"
// 	"go.mongodb.org/mongo-driver/mongo"
// 	"go.mongodb.org/mongo-driver/mongo/options"
// )

// // You will be using this Trainer type later in the program
// // type Trainer struct {
// // 	Name string
// // 	Age  int
// // 	City string
// // }

// func main() {

// 	// Rest of the code will go here
// 	// Set client options
// 	clientOptions := options.Client().ApplyURI("mongodb://localhost:2000")

// 	// Connect to MongoDB
// 	client, err := mongo.Connect(context.TODO(), clientOptions)

// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	// Check the connection
// 	err = client.Ping(context.TODO(), nil)

// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	fmt.Println("Connected to MongoDB!")
// 	// creates the collection
// 	collection := client.Database("Pokemon").Collection("trainers")

// 	// create a new trainer
// 	ash := Trainer{"Ash", 10, "Pallet Town"}
// 	misty := Trainer{"Mistry", 10, "Pallet Town"}

// 	// insert ash into the collection
// 	collection.InsertOne(context.TODO(), misty)
// 	collection.InsertOne(context.TODO(), ash)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	// fmt.Println("Inserted a single document: ", insertResult.InsertedID)
// 	func createTrainer(response http.ResponseWriter, request *http.Request) {
// 		response.Header().Set("content-type", "application/json")
// 		var person Person
// 		_ = json.NewDecoder(request.Body).Decode(&person)
// 		collection := client.Database("Trainers").Collection("seasonOne")
// 		ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
// 		result, _ := collection.InsertOne(ctx, person)
// 		json.NewEncoder(response).Encode(result)
// 	}

// 	router := mux.NewRouter()
// 	router.HandleFunc("/trainers", PostTrainer(ash)).Methods("POST")

// }
