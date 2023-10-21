package utils 

import (
	"fmt"
	"time"
	"os"
	"github.com/golang-jwt/jwt/v5"

	"github.com/starlingilcruz/golang-chat/internal/models"
)

func GenerateJWT(user models.User) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.RegisteredClaims{
		// A usual scenario is to set the expiration time relative to the current time
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
		IssuedAt:  jwt.NewNumericDate(time.Now()),
		NotBefore: jwt.NewNumericDate(time.Now()),
		Issuer:    os.Getenv("JWT_ISSUER"),
		Subject:   user.UserName,
		ID:        fmt.Sprint(user.ID),
		Audience:  []string{os.Getenv("JWT_AUDIENCE")},
	})
	
	tokenStr, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))

	return tokenStr, err
}