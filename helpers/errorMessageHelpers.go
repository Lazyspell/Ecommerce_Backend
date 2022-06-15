package helpers

import (
	"encoding/json"
	"net/http"
)

func ServerError(w http.ResponseWriter, err error) {
	w.WriteHeader(http.StatusInternalServerError)
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(err)
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

func NoContent204(w http.ResponseWriter, message string) {
	w.WriteHeader(http.StatusNoContent)
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(message)
}

func DeleteSuccessContent(w http.ResponseWriter) {
	w.WriteHeader(http.StatusOK)
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode("Success")
}

func UnAuthenticated(w http.ResponseWriter, message string) {
	w.WriteHeader(http.StatusUnauthorized)
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(message)
}
