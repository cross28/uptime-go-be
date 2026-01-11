package middleware

import "net/http"

func JwtAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		
		next.ServeHTTP(w, r)
	})
}