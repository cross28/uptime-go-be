package auth

import (
	"crosssystems.co/uptime-go-be/internal/auth/login"
	"github.com/go-chi/chi"
)

func RegisterRoutes(r chi.Router) {
	r.Post("/login", login.Login)
}