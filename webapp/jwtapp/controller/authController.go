package controller

import (
	"encoding/json"
	"fmt"
	"net/http"

	"golang.org/x/crypto/bcrypt"
	"pixels.com/jwtapp/config"
	"pixels.com/jwtapp/model"
	"pixels.com/jwtapp/utils"
)

func Register(w http.ResponseWriter, r *http.Request) {
	var user model.User
	json.NewDecoder(r.Body).Decode(&user)

	// Hash the password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		http.Error(w, "Error while hashing password", http.StatusInternalServerError)
		return
	}
	user.Password = string(hashedPassword)
	fmt.Println("db:", config.DBconfig.DB)
	config.DBconfig.DB.Create(&user)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}

func Login(w http.ResponseWriter, r *http.Request) {
	var user model.User
	var foundUser model.User
	json.NewDecoder(r.Body).Decode(&user)

	// Find user by username
	config.DBconfig.DB.Where("username = ?", user.Username).First(&foundUser)
	if foundUser.ID == 0 {
		http.Error(w, "User not found", http.StatusUnauthorized)
		return
	}

	// Check password
	err := bcrypt.CompareHashAndPassword([]byte(foundUser.Password), []byte(user.Password))
	if err != nil {
		http.Error(w, "Invalid password", http.StatusUnauthorized)
		return
	}

	// Generate JWT token
	token, err := utils.GenerateJWT(foundUser.Username)
	if err != nil {
		http.Error(w, "Error generating token", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"token": token})
}
