package util

import (
	"github.com/golang-jwt/jwt/v5"
)

type Header struct {
	Alg string `json:"alg"`
	Typ string `json:"typ"`
}

type Payload struct {
	Sub         int `json:"sub"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Email       string `json:"email"`
}

func CreateJWT(secretKey string, payload Payload) (string, error) {
	// jwt with package jwt
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub":          payload.Sub,
		"first_name":   payload.FirstName,
		"last_name":    payload.LastName,
		"email":        payload.Email,
	})	
	tokenString, err := token.SignedString([]byte(secretKey))
	return tokenString, err
}
