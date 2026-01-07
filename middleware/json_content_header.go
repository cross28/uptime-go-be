package middleware

import "net/http"

// Auto-set the content type header to be JSON if not set/overriden
func JsonContentType(next http.Handler) http.Handler{
	return http.HandlerFunc(func (w http.ResponseWriter, r *http.Request) {
		if (w.Header().Get("Content-Type") == "") {
			w.Header().Add("Content-Type", "application/json")
		}
		next.ServeHTTP(w, r)
	})
}