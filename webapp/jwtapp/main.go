package main

import (
	"fmt"
	"log"
	"net/http"

	"pixels.com/jwtapp/config"
	"pixels.com/jwtapp/routes"
)

func main() {

	config.Connect()

	// Register routes
	router := routes.RegisterRoutes()
	fmt.Println("Server is running on port 8000...")
	log.Fatal(http.ListenAndServe(":8000", router))

}
