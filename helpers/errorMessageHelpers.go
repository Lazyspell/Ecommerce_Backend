package helpers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"runtime/debug"
)

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

func badRequest400(w http.ResponseWriter, message string) {
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
