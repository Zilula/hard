package routes

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/Zilula/hard/db"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Api struct {
	db         *mongo.Database
	Collection *mongo.Collection
}

func New(db *mongo.Database) Api {
	return Api{
		db:         db,
		Collection: db.Collection("trainers"),
	}
}

// CREATE A NEW TRAINER
func (api Api) CreateTrainer() func(http.ResponseWriter, *http.Request) {

	// create route handler
	route := func(w http.ResponseWriter, r *http.Request) {
		// set content type
		w.Header().Set("Content-Type", "application/json")
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			panic(err)
		}
		// create new instance of our model
		var trainer db.Trainer

		// convert the JSON bytes and place them into the new trainer Struct
		json.Unmarshal([]byte(body), &trainer)

		// _ = json.NewDecoder(r.Body).Decode(&trainer)

		fmt.Println(trainer)
		// insert the body of the req into the DB
		api.Collection.InsertOne(context.TODO(), trainer)

		data, _ := json.Marshal(trainer)
		// import "encoding/json"
		//might not want to ignore error, might be ok
		w.Write(data)

	}

	// listen
	return route

}

type ResTrainer struct {
	ID   primitive.ObjectID `bson:"_id,omitempty"` // omitempty to protect against zeroed _id insertion
	Name string             `json:"Name" bson:"Name"`
	Age  []int              `json:"Age" bson:"Age"`
	City []string           `json:"City" bson:"City"`
}

// GET ALL TRAINERS
func (api Api) GetAllTrainers() func(http.ResponseWriter, *http.Request) {

	// create route handler
	route := func(w http.ResponseWriter, r *http.Request) {

		api.Collection.Find(context.Background(), bson.D{{}})

		fmt.Println("RESULTS", results)
		// data := json.Unmarshal([]byte(results))
		// // import "encoding/json"
		// //might not want to ignore error, might be ok
		// w.Write(data)

	}

	// listen
	return route

}
