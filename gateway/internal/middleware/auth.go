package middleware

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
)

type AuthMiddleware struct {
	AuthURL string
}

func NewAuthMiddleware(authURL string) *AuthMiddleware {
	return &AuthMiddleware{
		AuthURL: authURL,
	}
}

func (m *AuthMiddleware) Validate(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("Authorization")
		if token == "" {
			log.Printf("Authorization token is required")
			http.Error(w, "Authorization token is required", http.StatusUnauthorized)
			return
		}

		// создаем запрос для валидации токена
		reqBody, err := json.Marshal(map[string]string{
			"token": token,
		})
		if err != nil {
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}

		// отправляем запрос на сервис аутентификации
		resp, err := http.Post(m.AuthURL, "application/json", bytes.NewBuffer(reqBody))
		if err != nil {
			log.Printf("Error sending request to auth service: %v", err)
			http.Error(w, "Failed to validate token", http.StatusUnauthorized)
			return
		}
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusOK {
			log.Printf("Auth service returned status code: %d", resp.StatusCode)
			http.Error(w, "Invalid token", http.StatusUnauthorized)
			return
		}

		next.ServeHTTP(w, r)
	}
}
