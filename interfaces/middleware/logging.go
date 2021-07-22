package middleware

import (
	"net/http"

	log "github.com/sirupsen/logrus"
)

func (m *Middleware) Logging(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Infof("[%s] %s", r.Method, r.RequestURI)
		next.ServeHTTP(w, r)
	})
}
