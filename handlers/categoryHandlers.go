package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/lazyspell/Ecommerce_Backend/helpers"
	"github.com/lazyspell/Ecommerce_Backend/models"
)

func (m *Repository) GetAllCategories(w http.ResponseWriter, r *http.Request) {
	categories, err := m.DB.AllCategories()
	if err != nil {
		helpers.ServerError(w, err)
	}

	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(categories)
}

func (m *Repository) GetCategoryById(w http.ResponseWriter, r *http.Request) {
	var payload models.Categories
	err := json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		log.Println(err)
	}

	category, err := m.DB.CategoryById(payload.Id)
	if err != nil {
		helpers.ServerError(w, err)
	}

	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(category)
}
