package handlers

import (
	"fmt"
	"net/http"
)

type UserHandler struct{}

func NewUserHandler() *UserHandler {
	return &UserHandler{}
}
func (uh *UserHandler) GetAllUsers(rw http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(rw, "Get All users called")
}

func (uh *UserHandler) AddNewUser(rw http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(rw, "Add new user called")
}

func UserMiddleware(next http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("user middleware called")
		next.ServeHTTP(w, r)
	})
}
