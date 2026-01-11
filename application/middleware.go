package application

import (
	"crosssystems.co/uptime-go-be/middleware"
	"github.com/go-chi/chi"
	chi_middleware "github.com/go-chi/chi/middleware"
	"github.com/go-chi/cors"
)

func RegisterMiddleware(r *chi.Mux) {
	cors := cors.New(cors.Options{
		AllowedOrigins: []string{"https://*", "http://*"},
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTION"},
		AllowedHeaders: []string{"*"},
	})

	r.Use(chi_middleware.Logger)
	r.Use(chi_middleware.StripSlashes)
	r.Use(chi_middleware.RequestID)
	r.Use(chi_middleware.RealIP)
	r.Use(chi_middleware.Recoverer)
	r.Use(middleware.JsonContentType)
	r.Use(cors.Handler)
}
