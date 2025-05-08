package middleware

import (
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"strings"
)

var jwtSecret []byte("token")

func ValidateToken(token string) bool {
	if token == "" {
		return false
	}

	token = strings.TrimPrefix(token, "Bearer ")

	tokenParce, err := jwt.Parse(token, func (j *jwt.Token) (interface{}, error) {
		if _, ok := j.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid token")
		}
		return jwtSecret, nil
	})

	if err != nil {
		return false
	}

	if !tokenParce.Valid {
		return false
	}
	return true
}