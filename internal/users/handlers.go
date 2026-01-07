package users

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi"
)

// Similar to a primary constructor. This is the idiomatic way to make a constructor
type UserHandler struct {
	UserRepo UserRepository
}

func NewUserHandler(userRepo UserRepository) *UserHandler {
	return &UserHandler{
		UserRepo: userRepo,
	}
}

func (h *UserHandler) getUserById(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	user, err := h.UserRepo.GetById(id)
	if err != nil {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(user)
}

func (h *UserHandler) createUser(w http.ResponseWriter, r *http.Request) {
	var user User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	id, err := h.UserRepo.Create(&user)
	if err != nil {
		http.Error(w, "Failed to create user", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"id": id})
}