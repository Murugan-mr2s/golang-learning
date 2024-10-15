package main

import (
	"net/http"

	"github.com/gorilla/mux"
	"pixels.com/jsonv/handlers"
)

func main() {

	mx := mux.NewRouter()
	ph := handlers.NewProfileHandler()

	profileRouter := mx.PathPrefix("/profiles").Subrouter()
	profileRouter.HandleFunc("", ph.GetAllProfile).Methods(http.MethodGet)

	profileRouterPost := profileRouter.NewRoute().Subrouter()
	profileRouterPost.HandleFunc("", ph.AddNewProfile).Methods(http.MethodPost)
	profileRouterPost.Use(handlers.ProfileValidator)

	http.ListenAndServe(":8080", mx)
}
