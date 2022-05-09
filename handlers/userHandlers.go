package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"net/mail"

	"github.com/go-chi/jwtauth"
	"github.com/lazyspell/Ecommerce_Backend/helpers"
	"github.com/lazyspell/Ecommerce_Backend/models"
	"github.com/lazyspell/Ecommerce_Backend/utils"
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
	_, claims, err := jwtauth.FromContext(r.Context())
	if err != nil {
		log.Println(err)
		return
	}

	log.Println(claims["first_name"])
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

func (m *Repository) LoginUser(w http.ResponseWriter, r *http.Request) {
	var payload models.Users
	err := json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		helpers.BadRequest400(w, "Unable to decode request body please check request body")
		return
	}

	if payload.Email == "" {
		helpers.BadRequest400(w, "Email parameter not present in request body. check request body contents")
		return
	}

	utils.GenerateStateJwtCookie(w, payload)

	// if payload.Password == "" {
	// 	helpers.BadRequest400(w, "Password parameter not present in request body. check request body contents")
	// 	return
	// }

	// w.Header().Add("Content-Type", "application/json")
	// json.NewEncoder(w).Encode(payload)
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
