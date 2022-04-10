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
	errUnexpected   *usecase.UnexpectedError
	errNotFound     *usecase.NotFoundError
	errUnauthorized *usecase.UnauthorizedError
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
