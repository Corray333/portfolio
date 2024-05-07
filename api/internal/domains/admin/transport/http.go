package transport

import (
	"encoding/json"
	"io"
	"log/slog"
	"net/http"
	"os"

	"github.com/Corray333/portfolio/pkg/server/auth"
)

var refreshToken string = ""

// LogIn logs in the user and sends back refresh token, access token and user info
func LogIn() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		user := struct {
			Login    string `json:"login"`
			Password string `json:"password"`
		}{}
		body, err := io.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "Failed to read request body", http.StatusBadRequest)
			slog.Error("Failed to read request body: " + err.Error())
			return
		}
		if err := json.Unmarshal(body, &user); err != nil {
			http.Error(w, "Failed to unmarshal request body", http.StatusBadRequest)
			slog.Error("Failed to unmarshal request body: " + err.Error())
			return
		}

		if user.Login != os.Getenv("ADMIN") || user.Password != os.Getenv("PASSWORD") {
			http.Error(w, "Invalid credentials", http.StatusUnauthorized)
			slog.Error("Invalid credentials")
			return
		}

		refreshToken, err = auth.CreateToken(auth.RefreshTokenLifeTime)
		if err != nil {
			http.Error(w, "Failed to create token", http.StatusInternalServerError)
			slog.Error("Failed to create token: " + err.Error())
			return
		}

		token, err := auth.CreateToken(auth.AccessTokenLifeTime)
		if err != nil {
			http.Error(w, "Failed to create token", http.StatusInternalServerError)
			slog.Error("Failed to create token: " + err.Error())
			return
		}

		if err := json.NewEncoder(w).Encode(struct {
			Authorization string `json:"authorization"`
			Refresh       string `json:"refresh"`
		}{
			Authorization: token,
			Refresh:       refreshToken,
		}); err != nil {
			http.Error(w, "Failed to encode response", http.StatusInternalServerError)
			slog.Error("Failed to send response: " + err.Error())
			return
		}
	}
}

// RefreshAccessToken creates a new access token
func RefreshAccessToken() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		refresh := r.Header.Get("Refresh")
		var err error
		if refreshToken != refresh {
			http.Error(w, "Invalid refresh token", http.StatusUnauthorized)
			slog.Error("Invalid refresh token")
			return
		}
		refreshToken, err = auth.CreateToken(auth.RefreshTokenLifeTime)
		if err != nil {
			http.Error(w, "Failed to create token", http.StatusInternalServerError)
			slog.Error("Failed to create token: " + err.Error())
			return
		}

		access, err := auth.CreateToken(auth.AccessTokenLifeTime)
		if err != nil {
			http.Error(w, "Failed to create token", http.StatusInternalServerError)
			slog.Error("Failed to create token: " + err.Error())
			return
		}
		if err := json.NewEncoder(w).Encode(struct {
			Authorization string `json:"authorization"`
			Refresh       string `json:"refresh"`
		}{
			Authorization: access,
			Refresh:       refreshToken,
		}); err != nil {
			http.Error(w, "Failed to encode response", http.StatusInternalServerError)
			slog.Error("Failed to encode response: " + err.Error())
			return
		}
	}
}
