package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
	"pixels.com/userappdb/model"
)

type UserHandler struct {
	db *gorm.DB
}

func NewUserHandler(db *gorm.DB) *UserHandler {
	return &UserHandler{db}
}

func (uh *UserHandler) GetAllUsers(rw http.ResponseWriter, r *http.Request) {
	var users []model.User
	uh.db.Find(&users)
	rw.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(rw).Encode(users)
	if err != nil {
		http.Error(rw, "json marshel error", http.StatusInternalServerError)
	}
}

func (uh *UserHandler) AddNewUser(rw http.ResponseWriter, r *http.Request) {
	user := new(model.User)
	json.NewDecoder(r.Body).Decode(user)
	uh.db.Create(user)
	rw.Header().Set("Content-Type", "application/json")
	json.NewEncoder(rw).Encode(user)
}

func (uh *UserHandler) UpdateUser(rw http.ResponseWriter, r *http.Request) {
	user := new(model.User)

	json.NewDecoder(r.Body).Decode(user)
	uh.db.Save(user)

	//uh.db.Model(user).Update("name", user.Name).Update("email", user.Email)
	rw.Header().Set("Content-Type", "application/json")
	json.NewEncoder(rw).Encode(user)
}

func (uh *UserHandler) DeleteUser(rw http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]

	user := new(model.User)
	if err := uh.db.First(user, id).Error; err != nil {
		http.Error(rw, "user id not found", http.StatusBadRequest)
	}
	uh.db.Delete(user)
	rw.Header().Set("Content-Type", "application/json")
	json.NewEncoder(rw).Encode(map[string]string{"result": "success"})
}
