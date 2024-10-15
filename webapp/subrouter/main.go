package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"pixels.com/routes/handlers"
)

func HomeScreen(wr http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(wr, "Welcome sub router examples")
}

func GobalMiddleWare(next http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("gloal middle called")
		next.ServeHTTP(w, r)
	})
}

func main() {

	smx := mux.NewRouter()
	smx.HandleFunc("/", HomeScreen)
	smx.Use(GobalMiddleWare)

	uh := handlers.NewUserHandler()
	userRoute := smx.PathPrefix("/users").Subrouter()
	userRoute.Use(handlers.UserMiddleware)
	userRoute.HandleFunc("", uh.GetAllUsers).Methods(http.MethodGet)
	userRoute.HandleFunc("", uh.AddNewUser).Methods(http.MethodPost)

	ph := handlers.NewProductHandler()
	productRoute := smx.PathPrefix("/v1").Subrouter()
	productRoute.Use(handlers.ProductMiddleware)
	productRoute.HandleFunc("/products", ph.GetAllProducts).Methods(http.MethodGet)
	productRoute.HandleFunc("/products", ph.AddNewProduct).Methods(http.MethodPost)

	http.ListenAndServe(":8080", smx)

}
