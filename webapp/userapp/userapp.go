package main

import (
	"log"
	"net/http"
	"os"

	"examples.com/userapp/handlers"
	"github.com/gorilla/mux"
)

func main() {

	l := log.New(os.Stdout, "user-log", log.LstdFlags)
	router := mux.NewRouter()

	uh := handlers.NewUserHandler(l)
	router.HandleFunc("/users", uh.GetUsers).Methods(http.MethodGet)
	router.HandleFunc("/users/{id}", uh.GetUserById).Methods(http.MethodGet)
	router.HandleFunc("/users", uh.AddNewUser).Methods(http.MethodPost)
	router.HandleFunc("/users/{id}", uh.DeleteUserById).Methods(http.MethodDelete)

	http.ListenAndServe(":8080", router)

}
