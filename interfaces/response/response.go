package response

import (
	"encoding/json"
	"net/http"
)

func JSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	_, _ = w.Write(response)
}

func Error(w http.ResponseWriter, code int, message string) {
	JSON(w, code, map[string]string{
		"error": message,
	})
}

func OK(w http.ResponseWriter) {
	JSON(w, http.StatusOK, map[string]string{
		"message": "OK",
	})
}
