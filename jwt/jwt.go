package jwt

import (
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"

	"github.com/smmir-cent/Simple-Backend/handler"
)

type Config struct {
	AccessTokenSecret string
}

type JWT struct {
	Config
}

// NewAccessToken creates new access token for given user.
func (j JWT) NewAccessToken(u handler.User) (string, error) {
	// create token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		Audience:  "user",
		ExpiresAt: time.Now().Add(1 * time.Hour).Unix(),
		Id:        uuid.New().String(),
		IssuedAt:  time.Now().Unix(),
		Issuer:    "movie",
		NotBefore: time.Now().Unix(),
		Subject:   u.Email,
	})

	// generate encoded token and send it as response
	encodedToken, err := token.SignedString([]byte(j.AccessTokenSecret))
	if err != nil {
		return "", fmt.Errorf("failed to sign a token: %w", err)
	}

	return encodedToken, nil
}
