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

func (h *UserHandler) GetUserById(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	user, err := h.UserRepo.GetById(r.Context(), id)
	if err != nil {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(user)
}

func (h *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var user User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	id, err := h.UserRepo.Create(r.Context(), &user)
	if err != nil {
		http.Error(w, "Failed to create user", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"id": id})
}

func (h *UserHandler) GetAllUsers(w http.ResponseWriter, r *http.Request) {
	users, err := h.UserRepo.GetAllUsers(r.Context())
	if err != nil {
		http.Error(w, "Error getting all users", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(users)
}
