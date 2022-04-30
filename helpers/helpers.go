package helpers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"runtime/debug"

	"github.com/lazyspell/Ecommerce_Backend/config"
)

var app *config.AppConfig

func NewHelpers(a *config.AppConfig) {
	app = a
}

func ServerError(w http.ResponseWriter, err error) {
	trace := fmt.Sprintf("%s\n%s", err.Error(), debug.Stack())
	app.ErrorLog.Println(trace)
	http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
}

func Create201(w http.ResponseWriter) {
	w.WriteHeader(http.StatusCreated)
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode("Success")
}

func BadRequest400(w http.ResponseWriter, message string) {
	w.WriteHeader(http.StatusBadRequest)
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(message)
}
