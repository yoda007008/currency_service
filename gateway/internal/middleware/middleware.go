package middleware

import (
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"strings"
)

var jwtSecret = []byte("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0IiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxMjM0fQ.8oc-KZuuLCNwvkkpX2nFA4-YcDwcg2Y8e-WtvM--KSM")

func ValidateToken(token string) bool {
	if token == "" {
		return false
	}

	token = strings.TrimPrefix(token, "Bearer ")

	tokenParce, err := jwt.Parse(token, func(j *jwt.Token) (interface{}, error) {
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
