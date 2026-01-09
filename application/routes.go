package application

import (
	"crosssystems.co/uptime-go-be/internal/health"
	"crosssystems.co/uptime-go-be/internal/login"
	"crosssystems.co/uptime-go-be/internal/register"
	"crosssystems.co/uptime-go-be/internal/users"
	"github.com/go-chi/chi"
)

func (a *App) RegisterRoutes() {
	r := chi.NewRouter()

	RegisterMiddleware(r)

	r.Route("/auth", a.registerAuthRoutes)
	r.Route("/users", a.registerUserRoutes)
	r.Route("/health", a.registerHealthCheckRoute)

	a.Router = r
}

func (a *App) registerHealthCheckRoute(r chi.Router) {
	r.Get("/", health.Healthcheck)
}

func (a *App) registerUserRoutes(r chi.Router) {
	h := &users.UserHandler{
		UserRepo: users.NewPostgresUserRepository(a.PostgresDb),
	}

	r.Post("/", h.CreateUser)
	r.Get("/", h.GetAllUsers)
	r.Get("/{id}", h.GetUserById)
}

func (a *App) registerAuthRoutes(r chi.Router) {
	loginHandler := &login.LoginHandler{
		LoginRepo: login.NewPostgresLoginRepo(a.PostgresDb),
	}

	registerHandler := &register.RegisterHandler{
		RegistrationRepo: register.NewPostgresRegisterRepo(a.PostgresDb),
	}

	r.Post("/login", loginHandler.Login)
	r.Post("/register", registerHandler.Register)
}
