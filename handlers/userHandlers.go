package handlers

import (
	"encoding/json"
	"net/http"
	"net/mail"

	"github.com/lazyspell/Ecommerce_Backend/helpers"
	"github.com/lazyspell/Ecommerce_Backend/models"
	"golang.org/x/crypto/bcrypt"
)

func (m *Repository) NewUser(w http.ResponseWriter, r *http.Request) {
	var payload models.Users
	err := json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		helpers.BadRequest400(w, "invalid type please check request body")
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
		return
	}

	helpers.Create201(w)
}

func (m *Repository) GetAllUsers(w http.ResponseWriter, r *http.Request) {
	users, err := m.DB.AllUsers()
	if err != nil {
		helpers.ServerError(w, err)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)

}

func (m *Repository) GetUserById(w http.ResponseWriter, r *http.Request) {
	var payload models.Users
	err := json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		helpers.BadRequest400(w, "invalid type please check request body")
		return
	}

	if payload.Id == 0 {
		helpers.CheckValidId(w, payload)
		return
	}

	user, err := m.DB.UserById(payload.Id)
	if err != nil {
		helpers.ServerError(w, err)
		return
	}

	if user.Id == 0 {
		helpers.NoContent204(w, "No Document Retrived")
		return
	}

	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}

func (m *Repository) DeleteUser(w http.ResponseWriter, r *http.Request) {
	var payload models.Users
	err := json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		helpers.BadRequest400(w, "invalid type please check request body")
		return
	}

	_, err = m.DB.DeleteUserDB(payload.Id)
	if err != nil {

	}

	helpers.DeleteSuccessContent(w)
	return

}

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func validMailAddress(address string) bool {
	_, err := mail.ParseAddress(address)
	if err != nil {
		return false
	}
	return true
}
