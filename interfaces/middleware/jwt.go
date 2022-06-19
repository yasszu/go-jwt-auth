package middleware

import (
	"net/http"
	"strings"

	"github.com/yasszu/go-jwt-auth/interfaces/presenter"
)

func (h *Middleware) JWT() func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			authHeader := r.Header.Get("Authorization")
			extractedToken := strings.Split(authHeader, "Bearer ")
			if len(extractedToken) != 2 {
				presenter.Error(w, http.StatusForbidden, "Invalid header")
				return
			}

			token := strings.TrimSpace(extractedToken[1])
			accountID, err := h.jwtService.Verify(token)
			if err != nil {
				presenter.Error(w, http.StatusForbidden, "Invalid token")
				return
			}

			ctx := SetAccountID(r.Context(), accountID)
			next.ServeHTTP(w, r.WithContext(ctx))
			return
		})
	}
}
