package utils

import (
	"crypto/rand"
	"encoding/base64"
	"log"
	"net/http"
	"time"

	"github.com/lazyspell/Ecommerce_Backend/models"
)

func GenerateStateOauthCookie(w http.ResponseWriter) string {
	var expiration = time.Now().Add(2 * time.Minute)

	b := make([]byte, 16)
	rand.Read(b)
	state := base64.URLEncoding.EncodeToString(b)
	cookie := http.Cookie{
		Name:     "oauthstate",
		Value:    state,
		Expires:  expiration,
		HttpOnly: true,
	}
	http.SetCookie(w, &cookie)

	return state
}

func GenerateStateJwtCookie(w http.ResponseWriter, user models.Users) string {

	// var expiration = time.Now().Add(2 * time.Minute)

	state, err := jwtToken(user)
	if err != nil {
		log.Println(err)
		return state
	}

	cookie := http.Cookie{
		Name:  "jwt",
		Value: state,
		// Expires:  expiration,
		HttpOnly: true,
	}
	http.SetCookie(w, &cookie)

	return state

}

func GenerateGoogleJwtCookie(w http.ResponseWriter, user models.GoogleObject) string {

	var expiration = time.Now().Add(2 * time.Minute)

	state, err := googleJwtToken(user)
	if err != nil {
		log.Println(err)
		return state
	}

	cookie := http.Cookie{
		Name:     "jwt",
		Value:    state,
		Expires:  expiration,
		HttpOnly: true,
	}
	http.SetCookie(w, &cookie)

	return state

}
