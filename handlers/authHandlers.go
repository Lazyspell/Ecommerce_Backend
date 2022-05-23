package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/lazyspell/Ecommerce_Backend/helpers"
	"github.com/lazyspell/Ecommerce_Backend/models"
	"github.com/lazyspell/Ecommerce_Backend/utils"
	"golang.org/x/crypto/bcrypt"
)

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

	if !validMailAddress(payload.Email) {
		helpers.BadRequest400(w, "invalid email address given")
		return
	}

	if payload.Password == "" {
		helpers.BadRequest400(w, "Password parameter not present in request body. check request body contents")
		return
	}

	authUser, err := m.DB.Authenticate(payload.Email)
	if err != nil {
		helpers.BadRequest400(w, "No User Found")
		return
	}

	if !compareHashPassword(payload.Password, []byte(authUser.Password)) {
		helpers.UnAuthenticated(w, "invalid password given")
		return
	}

	utils.GenerateStateJwtCookie(w, authUser)

	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode("")
}

//not working properly need to fix
func compareHashPassword(password string, hashedPassword []byte) bool {
	err := bcrypt.CompareHashAndPassword(hashedPassword, []byte(password))
	if err != nil {
		log.Println("passwords did not match. unauthenticated")
		return false
	}

	return true
}
