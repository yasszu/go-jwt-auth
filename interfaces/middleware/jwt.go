package middleware

import (
	"net/http"
	"strings"

	"github.com/yasszu/go-jwt-auth/infrastructure/jwt"
	"github.com/yasszu/go-jwt-auth/interfaces/presenter"
)

func JWT(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		extractedToken := strings.Split(authHeader, "Bearer ")
		if len(extractedToken) != 2 {
			presenter.Error(w, http.StatusForbidden, "Forbidden")
			return
		}

		token := strings.TrimSpace(extractedToken[1])
		claims, err := jwt.ValidateToken(token)
		if err != nil {
			presenter.Error(w, http.StatusForbidden, "Forbidden")
			return
		}

		ctx := SetAccountID(r.Context(), claims.AccountID)
		next.ServeHTTP(w, r.WithContext(ctx))
		return
	})
}
