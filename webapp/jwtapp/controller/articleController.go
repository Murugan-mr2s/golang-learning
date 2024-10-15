package controller

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"pixels.com/jwtapp/config"
	"pixels.com/jwtapp/model"
)

// Get all articles
func GetArticles(w http.ResponseWriter, r *http.Request) {
	var articles []model.Article
	config.DBconfig.DB.Find(&articles)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(articles)
}

// Create a new article
func CreateArticle(w http.ResponseWriter, r *http.Request) {
	var article model.Article
	json.NewDecoder(r.Body).Decode(&article)
	config.DBconfig.DB.Create(&article)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(article)
}

// Get a single article by ID
func GetArticle(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var article model.Article
	config.DBconfig.DB.First(&article, params["id"])
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(article)
}

// Update an article by ID
func UpdateArticle(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var article model.Article
	config.DBconfig.DB.First(&article, params["id"])
	json.NewDecoder(r.Body).Decode(&article)
	config.DBconfig.DB.Save(&article)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(article)
}

// Delete an article by ID
func DeleteArticle(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var article model.Article
	config.DBconfig.DB.First(&article, params["id"])
	config.DBconfig.DB.Delete(&article)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"result": "success"})
}
