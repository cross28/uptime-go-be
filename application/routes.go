package application

import (
	"crosssystems.co/uptime-go-be/internal/auth"
	"crosssystems.co/uptime-go-be/internal/health"
	"crosssystems.co/uptime-go-be/internal/users"
	"github.com/go-chi/chi"
)

func RegisterRoutes() *chi.Mux {
	r := chi.NewRouter()

	RegisterMiddleware(r)
	
	r.Route("/", auth.RegisterRoutes)
	r.Route("/user", users.RegisterRoutes)
	r.Route("/health", health.RegisterRoutes)

	return r
}