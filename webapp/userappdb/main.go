package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"pixels.com/userappdb/handlers"
	"pixels.com/userappdb/repo"
)

func main() {

	l := log.New(os.Stdout, "user-app-db", log.LstdFlags)
	mx := mux.NewRouter()

	db := repo.InitDB()
	uh := handlers.NewUserHandler(db)

	mx.HandleFunc("/users", uh.GetAllUsers).Methods(http.MethodGet)
	mx.HandleFunc("/users", uh.AddNewUser).Methods(http.MethodPost)
	mx.HandleFunc("/users", uh.UpdateUser).Methods(http.MethodPut)
	mx.HandleFunc("/users/{id:[0-9+]}", uh.DeleteUser).Methods(http.MethodDelete)

	err := http.ListenAndServe(":8080", mx)
	if err != nil {
		l.Fatal("server can not start")
	}
	fmt.Println("Server Starts listening... ")
}
