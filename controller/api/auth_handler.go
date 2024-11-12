package api

import (
	"encoding/json"
	"net/http"

	"github.com/Muhfikri12/project-app-inventaris-golang-fikri/model"
	apiservice "github.com/Muhfikri12/project-app-inventaris-golang-fikri/service/api_service"
	"github.com/Muhfikri12/project-app-inventaris-golang-fikri/utils/validation"
)

type AuthHandler struct {
	AuthService *apiservice.AuthService
}

func NewAuthHandler(service *apiservice.AuthService) *AuthHandler {
	return &AuthHandler{AuthService: service}
}

func (ah *AuthHandler) LoginHandler(w http.ResponseWriter, r *http.Request) {
	var user model.User

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		validation.BadResponse(w, "Invalid request payload" + err.Error(), http.StatusBadRequest)
		return
	}

	err = ah.AuthService.LoginUser(&user)
	if err != nil {
		validation.BadResponse(w, "Data tidak tersedia", http.StatusNotFound)
		return
	}

	validation.OKResponse(w, "successfully login", user.Token)
}

func (ah *AuthHandler) Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Mengambil token dari header kustom "Token"
		token := r.Header.Get("Authorization")

		// Cek apakah token ada
		if token == "" {
			validation.BadResponse(w, "header token tidak ditemukan", http.StatusUnauthorized)
			return
		}

		// Validasi token
		err := ah.AuthService.CheckToken(token)
		if err != nil {
			validation.BadResponse(w, err.Error(), http.StatusUnauthorized)
			return
		}

		next.ServeHTTP(w, r)
	})
}

