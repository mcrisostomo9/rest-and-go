package main

import (
	"os"
	"log"
	"github.com/gorilla/handlers"
	"net/http"
	"rest-and-go/store"
)

func main() {
	//get the PORT env variable
	port := os.Getenv("PORT")

	if port == ""{
		log.Fatal("$PORT must be set")
	}

	// create routes
	router := store.NewRouter()

	// These two lines are important if you're designing a front-end to utilise this API methods
	allowedOrigins := handlers.AllowedOrigins([]string{"*"})
	allowedMethods := handlers.AllowedMethods([]string{"GET", "POST", "DELETE", "PUT"})

	//Launch server with CORS validations
	log.Fatal(http.ListenAndServe(":" + port, handlers.CORS(allowedOrigins, allowedMethods)(router)))
}