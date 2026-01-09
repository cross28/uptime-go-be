package register

import (
	"encoding/json"
	"net/http"

	"crosssystems.co/uptime-go-be/auth"
)

type RegisterHandler struct {
	RegistrationRepo RegisterRepo
}

func (h *RegisterHandler) Register(w http.ResponseWriter, r *http.Request) {
	var registerReq RegisterRequest
	err := json.NewDecoder(r.Body).Decode(&registerReq)
	if err != nil {
		http.Error(w, "error reading register request", http.StatusInternalServerError)
		return
	}

	password_hash, err := auth.HashPassword(registerReq.Password)
	if err != nil {
		http.Error(w, "error hashing password", http.StatusBadRequest)
		return
	}

	err = h.RegistrationRepo.Register(r.Context(), registerReq.Email, password_hash)
	if err != nil {
		http.Error(w, "error occurred registering user", http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
}