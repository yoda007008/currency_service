package middleware

import (
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"log"
)

var jwtSecret = []byte("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNzQ2OTUyNTg0fQ.7KvKNt4nQqSo0u9vTPssMAn_56onpZfrAxo5pUAng4g")

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
