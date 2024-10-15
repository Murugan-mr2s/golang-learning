package handlers

import (
	"fmt"
	"net/http"
)

type ProductHandler struct{}

func NewProductHandler() *ProductHandler {
	return &ProductHandler{}
}

func (ph *ProductHandler) GetAllProducts(rw http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(rw, "Get All products called")
}

func (ph *ProductHandler) AddNewProduct(rw http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(rw, "Add new Product called")
}

func ProductMiddleware(next http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("product middleware called")
		next.ServeHTTP(w, r)
	})
}

func ProductMethodMiddleware(next http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("product specific method middleware called")
		next.ServeHTTP(w, r)
	})
}
