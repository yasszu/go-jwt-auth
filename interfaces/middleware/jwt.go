package middleware

import (
	"context"
	"go-jwt-auth/domain/entity"
	"go-jwt-auth/infrastructure/jwt"
	"go-jwt-auth/interfaces/response"
	"net/http"
	"strings"
)

func (m *Middleware) JWT(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		extractedToken := strings.Split(authHeader, "Bearer ")
		if len(extractedToken) == 2 {
			token := strings.TrimSpace(extractedToken[1])

			claims, err := jwt.ValidateToken(token)
			if err != nil {
				response.Error(w, http.StatusForbidden, "Forbidden")
				return
			}

			ctx := context.WithValue(r.Context(), entity.ContextKeyAccountID, claims.AccountID)
			next.ServeHTTP(w, r.WithContext(ctx))
			return
		} else {
			response.Error(w, http.StatusForbidden, "Forbidden")
			return
		}
	})
}
