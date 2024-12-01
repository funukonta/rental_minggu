package auth

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type Auth struct {
	Token string `json:"token"`
}

func GenerateJWT(username, role string) (*Auth, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": username,
		"role":     role,
		"expires":  time.Now().Add(time.Hour * 24).Unix(),
	})

	// Sign and get the complete encoded token as a string using the secret
	secret := os.Getenv("SECRET_KEY")
	tokenString, err := token.SignedString([]byte(secret))
	if err != nil {
		return nil, err
	}

	auth := &Auth{
		Token: tokenString,
	}

	return auth, nil
}
