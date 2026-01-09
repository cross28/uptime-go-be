package login

import (
	"encoding/json"
	"errors"
	"net/http"

	"crosssystems.co/uptime-go-be/auth"
	"github.com/jackc/pgx/v5"
)

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginHandler struct {
	LoginRepo LoginRepo
}

func (h *LoginHandler) Login(w http.ResponseWriter, r *http.Request) {
	var login LoginRequest
	if err := json.NewDecoder(r.Body).Decode(&login); err != nil {
		http.Error(w, "Invalid login request body", http.StatusBadRequest)
		return
	}

	password_hash, err := h.LoginRepo.GetPasswordHash(r.Context(), login.Email)

	if errors.Is(err, pgx.ErrNoRows) {
		http.Error(w, "Email not found", http.StatusBadRequest)
		return
	} else if err != nil {
		http.Error(w, "Error creating login", http.StatusInternalServerError)
		return
	}

	match := auth.VerifyPassword(login.Password, password_hash)
	if !match {
		http.Error(w, "Email or password incorrect", http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
}
