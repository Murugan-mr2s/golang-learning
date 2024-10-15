package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func IOReadHandler(rw http.ResponseWriter, r *http.Request) {

	Vars := mux.Vars(r)
	articleID := Vars["id"]

	comments := r.URL.Query().Get("comments")

	fmt.Fprintln(rw, "article Id:", articleID)
	fmt.Fprintln(rw, "article comments:", comments)
}

func IOReadArrayHandler(rw http.ResponseWriter, r *http.Request) {

	Vars := mux.Vars(r)
	articleID := Vars["id"]

	categoryList := r.URL.Query()["category"]

	fmt.Fprintln(rw, "article Id:", articleID)
	fmt.Fprintln(rw, "article comments:", categoryList)
}

func main() {

	router := mux.NewRouter()

	router.HandleFunc("/articles/{id}", IOReadHandler).Methods("GET")
	router.HandleFunc("/articles/{id}/categories", IOReadArrayHandler).Methods("GET")

	fmt.Println("Server starting on 8080")
	log.Fatal(http.ListenAndServe(":8080", router))

}
