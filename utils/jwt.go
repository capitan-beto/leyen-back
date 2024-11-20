package utils

import (
	"fmt"
	"os"
	s "strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var secret = []byte(os.Getenv("JWT_SECRET"))

func GenerateToken(email, role string) (*string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email": email,
		"role":  role,
		"exp":   time.Now().Add(time.Hour * 24).Unix(),
	})

	tokenString, err := token.SignedString(secret)
	if err != nil {
		return nil, nil
	}

	return &tokenString, nil
}

func VerifyToken(tokenString, role string) error {
	token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		return secret, nil
	})
	if err != nil {
		return err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		if !s.Contains(claims["role"].(string), role) {
			return fmt.Errorf("permission denied")
		}
	} else {
		return fmt.Errorf("invalid token")
	}

	return err
}
