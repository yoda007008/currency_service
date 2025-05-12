package middleware

import (
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"log"
)

var jwtSecret = []byte("a-string-secret-at-least-256-bits-long")

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
