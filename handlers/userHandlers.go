package handlers

// import (
// 	"encoding/json"
// 	"log"
// 	"net/http"

// 	"github.com/lazyspell/Ecommerce_Backend/models"
// )

// func (m *Repository) NewUser(w http.ResponseWriter, r *http.Request) string {
// 	var payload models.Users
// 	err := json.NewDecoder(r.Body).Decode(&payload)
// 	if err != nil {
// 		log.Println(err)
// 	}

// 	var user models.Users

// 	user.FirstName = payload.FirstName
// 	user.LastName = payload.LastName
// 	user.Email = payload.Email
// 	user.Password = payload.Password

// 	_, err = m.DB.NewUserDB(payload)
// 	if err != nil {
// 		log.Println(err)
// 	}

// 	return "Success"
// }
