package routes

import (
	"github.com/gorilla/mux"
	"pixels.com/jwtapp/controller"
	"pixels.com/jwtapp/middleware"
)

func RegisterRoutes() *mux.Router {
	router := mux.NewRouter()

	// Authentication routes
	router.HandleFunc("/register", controller.Register).Methods("POST")
	router.HandleFunc("/login", controller.Login).Methods("POST")

	// Articles routes (protected by JWT)
	articleRoutes := router.PathPrefix("/articles").Subrouter()
	articleRoutes.Use(middleware.JWTMiddleware)
	articleRoutes.HandleFunc("", controller.GetArticles).Methods("GET")
	articleRoutes.HandleFunc("", controller.CreateArticle).Methods("POST")
	articleRoutes.HandleFunc("/{id}", controller.GetArticle).Methods("GET")
	articleRoutes.HandleFunc("/{id}", controller.UpdateArticle).Methods("PUT")
	articleRoutes.HandleFunc("/{id}", controller.DeleteArticle).Methods("DELETE")

	return router
}
