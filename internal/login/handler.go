package login

import (
	"encoding/json"
	"errors"
	"net/http"
	"time"

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

	user_login, err := h.LoginRepo.GetUserByEmail(r.Context(), login.Email)

	if errors.Is(err, pgx.ErrNoRows) {
		http.Error(w, "Email not found", http.StatusBadRequest)
		return
	} else if err != nil {
		http.Error(w, "Error creating login", http.StatusInternalServerError)
		return
	}

	match := auth.VerifyPassword(login.Password, user_login.PasswordHash)
	if !match {
		http.Error(w, "Email or password incorrect", http.StatusUnauthorized)
		return
	}

	access_token, err := auth.CreateJwtToken(user_login.Id)
	if err != nil {
		http.Error(w, "Error creating jwt", http.StatusInternalServerError)
		return
	}

	refresh_token, err := auth.CreateRefreshToken(user_login.Id)
	if err != nil {
		http.Error(w, "error creating refresh jwt", http.StatusInternalServerError)
		return
	}

	access_token_cookie := &http.Cookie{
		Name: "access_token",
		Value: access_token,
		Expires: time.Now().Add(time.Hour * 1),
		HttpOnly: true,
		Secure: true,
		Path: "/",
		SameSite: http.SameSiteLaxMode,
	}

	refresh_token_cookie := &http.Cookie{
		Name: "refresh_token",
		Value: refresh_token,
		Expires: time.Now().Add(time.Hour * 24 * 7),
		HttpOnly: true,
		Secure: true,
		Path: "/",
		SameSite: http.SameSiteLaxMode,
	}

	http.SetCookie(w, access_token_cookie)
	http.SetCookie(w, refresh_token_cookie)

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Login success"))
}
