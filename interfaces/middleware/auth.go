package middleware

import (
	"context"
	"go-jwt-auth/infrastructure/auth"
	"go-jwt-auth/interfaces/response"
	"net/http"
	"strings"
)

func JwtMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		extractedToken := strings.Split(authHeader, "Bearer ")
		if len(extractedToken) == 2 {
			token := strings.TrimSpace(extractedToken[1])

			claims, err := auth.ValidateToken(token)
			if err != nil {
				response.Error(w, http.StatusForbidden, "Forbidden")
				return
			}

			ctx := context.WithValue(r.Context(), auth.AccountIdKey, claims.AccountID)
			next.ServeHTTP(w, r.WithContext(ctx))
			return
		} else {
			response.Error(w, http.StatusForbidden, "Forbidden")
			return
		}
	})
}
