package presenter

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/yasszu/go-jwt-auth/application/usecase"
)

func JSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	_, _ = w.Write(response)
}

func NewError(w http.ResponseWriter, err error) {
	Error(w, Status(err), Message(err))
}

func NewBadRequest(w http.ResponseWriter) {
	Error(w, http.StatusBadRequest, "Bad Request")
}

func NewUnauthorized(w http.ResponseWriter) {
	Error(w, http.StatusUnauthorized, "Unauthorized")
}

func NewInternalServerError(w http.ResponseWriter) {
	Error(w, http.StatusInternalServerError, "Server Error")
}

func Error(w http.ResponseWriter, code int, message interface{}) {
	JSON(w, code, map[string]string{
		"error": fmt.Sprint(message),
	})
}

func OK(w http.ResponseWriter) {
	JSON(w, http.StatusOK, map[string]string{
		"message": "OK",
	})
}

var (
	errUnexpected   *usecase.ErrorUnexpected
	errNotFound     *usecase.ErrorNotFound
	errUnauthorized *usecase.ErrorUnauthorized
)

func Status(err error) int {
	switch {
	case errors.As(err, &errUnexpected):
		return http.StatusInternalServerError
	case errors.As(err, &errNotFound):
		return http.StatusNotFound
	case errors.As(err, &errUnauthorized):
		return http.StatusUnauthorized
	default:
		return http.StatusInternalServerError
	}
}

func Message(err error) string {
	switch {
	case errors.As(err, &errUnexpected):
		return "Server Error"
	case errors.As(err, &errNotFound):
		return "Not Found"
	case errors.As(err, &errUnauthorized):
		return "Unauthorized"
	default:
		return "Server Error"
	}
}
