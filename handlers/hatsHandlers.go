package handlers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/lazyspell/Ecommerce_Backend/helpers"
	"github.com/lazyspell/Ecommerce_Backend/models"
)

func (m *Repository) NewHats(w http.ResponseWriter, r *http.Request) {
	var payload models.Hats

	err := json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		helpers.BadRequest400(w, "invalid type please check request body")
		return
	}

	var hats models.Hats

	hats.HatName = payload.HatName
	hats.ImageUrl = payload.ImageUrl
	hats.Price = payload.Price
	hats.CreatedAt = time.Now()
	hats.UpdatedAt = time.Now()
	_, err = m.DB.NewHatsDB(hats)
	if err != nil {
		helpers.ServerError(w, err)
		return
	}

	helpers.Create201(w)

}
