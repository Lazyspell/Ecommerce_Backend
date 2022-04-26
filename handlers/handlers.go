package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/lazyspell/Ecommerce_Backend/config"
	"github.com/lazyspell/Ecommerce_Backend/driver"
	"github.com/lazyspell/Ecommerce_Backend/repository"
	"github.com/lazyspell/Ecommerce_Backend/repository/dbrepo"
)

var Repo *Repository

type Repository struct {
	App *config.AppConfig
	DB  repository.DatabaseRepo
}

func NewRepo(a *config.AppConfig, db *driver.DB) *Repository {
	return &Repository{
		App: a,
		DB:  dbrepo.NewPostgresRepo(db.SQL, a),
	}
}

func NewHandler(r *Repository) {
	Repo = r
}

func (m *Repository) GetAllCategories(w http.ResponseWriter, r *http.Request) {
	categories, err := m.DB.AllCategories()
	if err != nil {
		log.Println(err)
	}

	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(categories)
}
