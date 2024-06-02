package utils

import (
	"fmt"

	"github.com/golang-jwt/jwt"
)

type Jwt struct{}

type IJwt interface {
	SignToken() string
	VerifyToken() bool
}

type CustomClaims struct {
	Id   uint `json:"id"`
	Role int  `json:"role"`
}

func (c *CustomClaims) Valid() error {
	return nil
}

func NewCustomClaims(id uint, role int) *CustomClaims {
	return &CustomClaims{
		Id:   id,
		Role: role,
	}
}

func SignDocument(id uint, role int) string {
	claims := NewCustomClaims(id, role)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString([]byte("secret123"))
	if err != nil {
		fmt.Println("Error signing token: ")
		panic(err)
	}
	return signedToken
}

func VerifyToken(tokenString string) (map[string]interface{}, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte("secret123"), nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, fmt.Errorf("invalid token")
}
