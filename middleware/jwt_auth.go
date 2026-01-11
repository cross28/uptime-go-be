package middleware

import (
	"net/http"

	"crosssystems.co/uptime-go-be/auth"
)

func JwtAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie("access_token")
		if err != nil {
			if err == http.ErrNoCookie {
				http.Error(w, "Unauthorized", http.StatusUnauthorized)
				return
			}
			http.Error(w, "Bad request", http.StatusBadRequest)
			return
		}

		token := cookie.Value

		if auth.VerifyJwtToken(token) != nil {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}
		
		// access granted
		next.ServeHTTP(w, r)
	})
}