package users

import "github.com/go-chi/chi"

func RegisterRoutes(r chi.Router) {
	h := &UserHandler{
		UserRepo: NewPostgresUserRepository(nil),
	}

	r.Get("/{id}", h.getUserById)
	r.Post("/", h.createUser)
}