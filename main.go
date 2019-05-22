package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Zilula/hard/db"
	"github.com/Zilula/hard/routes"
)

func main() {
	// create a DB client 
	client := db.CreateConnection()

	// return a router
	router := routes.CreateRouter(client)

	log.Fatal(http.ListenAndServe(":8000", router))
	fmt.Println("Server listening on port 8000")
}
