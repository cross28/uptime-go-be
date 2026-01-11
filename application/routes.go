package application

import (
	"crosssystems.co/uptime-go-be/internal/health"
	"crosssystems.co/uptime-go-be/middleware"
	"github.com/go-chi/chi"
)

func (a *App) RegisterRoutes() {
	r := chi.NewRouter()

	RegisterMiddleware(r)

	r.Route("/auth", func(r chi.Router){
		r.Post("/login", a.LoginHandler.Login)
		r.Post("/register", a.RegisterHandler.Register)
	})

	r.Route("/users", func(r chi.Router){
		r.Use(middleware.JwtAuth)
		r.Post("/", a.UserHandler.CreateUser)
		r.Get("/", a.UserHandler.GetAllUsers)
		r.Get("/{id}", a.UserHandler.GetUserById)
	})

	r.Get("/health", health.Healthcheck)

	a.Router = r
}
