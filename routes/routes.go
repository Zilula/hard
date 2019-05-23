package routes

import (
	"fmt"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/mongo"
)

func CreateRouter(db *mongo.Database) (*mux.Router) {
	fmt.Println("db", db)
	// make a new router
	router := mux.NewRouter()

	// get a new connection to the DB/API
	api := New(db)

	// Route handlers // endpoints
	router.HandleFunc("/api/trainers", api.CreateTrainer()).Methods("POST")
	router.HandleFunc("/api/trainers", api.GetAllTrainers()).Methods("GET")

	//run server

	return router
}
