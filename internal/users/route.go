package users

import "github.com/go-chi/chi"

func RegisterRoutes(r chi.Router, h *UserHandler) {
	r.Get("/{id}", h.getUserById)
	r.Post("/", h.createUser)
}