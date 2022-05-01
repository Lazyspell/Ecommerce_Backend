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
		return
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
