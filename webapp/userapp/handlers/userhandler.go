package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"examples.com/userapp/model"
	"github.com/gorilla/mux"
)

type UserHandler struct {
	l *log.Logger
}

func NewUserHandler(l *log.Logger) *UserHandler {
	return &UserHandler{l}
}

func (u *UserHandler) GetUsers(rw http.ResponseWriter, r *http.Request) {

	rw.Header().Set("Content-Type", "application/json")
	users := model.GetUsers()
	err := users.ToJSON(rw)

	if err != nil {
		http.Error(rw, "unable to json marshal", http.StatusInternalServerError)
		return
	}
}

func (u *UserHandler) GetUserById(rw http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])

	if err != nil {
		http.Error(rw, "user id not found", http.StatusBadRequest)
		return
	}

	user, err := model.GetUserById(id)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusBadRequest)
		return
	}

	bs, err := json.Marshal(user)
	rw.Header().Set("Content-Type", "application/json")
	rw.Write(bs)

	if err != nil {
		http.Error(rw, "unable to json marshal", http.StatusInternalServerError)
		return
	}
}

func (u *UserHandler) AddNewUser(rw http.ResponseWriter, r *http.Request) {
	var user model.User
	err := json.NewDecoder(r.Body).Decode(&user)

	if err != nil {
		http.Error(rw, "unable to json marshel", http.StatusBadRequest)
		return
	}
	model.AddUser(&user)
	fmt.Fprintln(rw, "New User Added")
}

func (u *UserHandler) DeleteUserById(rw http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(rw, "user id not found", http.StatusBadRequest)
		return
	}

	err = model.DeleteUserById(id)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusBadRequest)
		return
	}
	fmt.Fprintln(rw, "User Deleted Successfully")
}
