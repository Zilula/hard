package routes

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/Zilula/hard/db"
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

func (api Api) CreateTrainer() func(http.ResponseWriter, *http.Request) {

	// create route handler
	route := func(w http.ResponseWriter, r *http.Request) {
		// set content type
		w.Header().Set("Content-Type", "application/json")
		body, err := ioutil.ReadAll(r.Body)

		fmt.Println("BODY", body, err)

		var trainer db.Trainer

		fmt.Println("BODY FROM RETURNED FUNC", r.Body)
		_ = json.NewDecoder(r.Body).Decode(&trainer)

		// insert the body of the req into the DB
		api.Collection.InsertOne(context.TODO(), r.Body)
	}

	// listen
	return route

}
