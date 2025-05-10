package middleware

import (
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"log"
)

var jwtSecret = []byte("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiYWRtaW4iOnRydWUsImlhdCI6MTc0NjkwMjMwOH0.C8EVxabGNSAN4Z7ErGRM99UNBZ2dlbOi7J50m6BQeW4")

func ValidateToken(token string) bool {
	if token == "" {
		log.Println("Token is empty")
		return false
	}

	// убираем префикс "Bearer "
	if len(token) <= len("Bearer ") || token[:len("Bearer ")] != "Bearer " {
		log.Println("Token format is incorrect")
		return false
	}

	token = token[len("Bearer "):] // Убираем префикс

	tokenParsed, err := jwt.Parse(token, func(j *jwt.Token) (interface{}, error) {
		if _, ok := j.Method.(*jwt.SigningMethodHMAC); !ok {
			log.Println("Invalid token signing method")
			return nil, errors.New("invalid token signing method")
		}
		return jwtSecret, nil
	})

	if err != nil {
		log.Printf("Error parsing token: %v", err)
		return false
	}

	// проверка валидности токена
	if !tokenParsed.Valid {
		log.Println("Token is not valid")
		return false
	}

	return true
}
