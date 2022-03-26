package middleware

import (
	"context"
	"net/http"
	"strings"

	"github.com/yasszu/go-jwt-auth/interfaces/presenter"

	"github.com/yasszu/go-jwt-auth/domain/entity"
	"github.com/yasszu/go-jwt-auth/infrastructure/jwt"
)

func (m *Middleware) JWT(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		extractedToken := strings.Split(authHeader, "Bearer ")
		if len(extractedToken) == 2 {
			token := strings.TrimSpace(extractedToken[1])

			claims, err := jwt.ValidateToken(token)
			if err != nil {
				presenter.Error(w, http.StatusForbidden, "Forbidden")
				return
			}

			ctx := context.WithValue(r.Context(), entity.ContextKeyAccountID, claims.AccountID)
			next.ServeHTTP(w, r.WithContext(ctx))
			return
		} else {
			presenter.Error(w, http.StatusForbidden, "Forbidden")
			return
		}
	})
}
