package response

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/yasszu/go-jwt-auth/domain/entity"
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
	errUnexpected   *entity.UnexpectedError
	errNotFound     *entity.NotFoundError
	errUnauthorized *entity.UnauthorizedError
)

func Status(err error) int {
	if errors.As(err, &errUnexpected) {
		return http.StatusInternalServerError
	}
	if errors.As(err, &errNotFound) {
		return http.StatusNotFound
	}
	if errors.As(err, &errUnauthorized) {
		return http.StatusUnauthorized
	}
	return http.StatusInternalServerError
}
