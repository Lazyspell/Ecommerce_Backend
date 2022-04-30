package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"net/mail"

	"github.com/lazyspell/Ecommerce_Backend/config"
	"github.com/lazyspell/Ecommerce_Backend/driver"
	"github.com/lazyspell/Ecommerce_Backend/helpers"
	"github.com/lazyspell/Ecommerce_Backend/models"
	"github.com/lazyspell/Ecommerce_Backend/repository"
	"github.com/lazyspell/Ecommerce_Backend/repository/dbrepo"
	"golang.org/x/crypto/bcrypt"
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

func (m *Repository) NewUser(w http.ResponseWriter, r *http.Request) {
	var payload models.Users
	err := json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		helpers.BadRequest400(w, "Bad Request After Getting Body")
		return
	}

	if !validMailAddress(payload.Email) {
		helpers.BadRequest400(w, "Invalid Email Address Given")
		return
	}

	var user models.Users

	user.FirstName = payload.FirstName
	user.LastName = payload.LastName
	user.Email = payload.Email
	user.Password, _ = hashPassword(payload.Password)
	_, err = m.DB.NewUserDB(user)
	if err != nil {
		helpers.ServerError(w, err)
	}

	helpers.Create201(w)
}

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func checkPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
func validMailAddress(address string) bool {
	_, err := mail.ParseAddress(address)
	if err != nil {
		return false
	}
	return true
}
