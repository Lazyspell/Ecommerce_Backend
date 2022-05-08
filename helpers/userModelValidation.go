package helpers

import (
	"net/http"

	"github.com/lazyspell/Ecommerce_Backend/models"
)

func CheckValidId(w http.ResponseWriter, payload models.Users) bool {
	if payload.Id == 0 {
		BadRequest400(w, "ID parameter not present in request body. check request body contents")
		return false
	}
	return true
}

func CheckValidPayload(err error, w http.ResponseWriter, payload models.Users) bool {
	if err != nil {
		BadRequest400(w, "invalid type please check request body")
		return false
	}
	return true
}
