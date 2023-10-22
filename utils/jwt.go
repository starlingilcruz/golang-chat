package utils 

import (
	"fmt"
	"log"
	"time"
	"os"
	"github.com/golang-jwt/jwt/v5"

	"github.com/starlingilcruz/golang-chat/internal/models"
)

type TokenClaims struct {
	jwt.RegisteredClaims
	UserId    uint   `json:"userid"`
	UserName  string `json:"username"`
	Email     string `json:"email"`
}

func GenerateJWT(user models.User) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, TokenClaims{
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
			Issuer:    os.Getenv("JWT_ISSUER"),
			Subject:   user.UserName,
			ID:        fmt.Sprint(user.ID),
			Audience:  []string{os.Getenv("JWT_AUDIENCE")},
		},
		user.ID,
		user.UserName,
		user.Email, 
	})

	return token.SignedString([]byte(os.Getenv("JWT_SECRET")))
}

func VerifyJWT(token string) (jwt.MapClaims, error) {
		parsedJWT, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return []byte(os.Getenv("JWT_SECRET")), nil
		})

		if err != nil {
			log.Fatalf("Couldn't parse token: %v", err)
		}

		return parsedJWT.Claims.(jwt.MapClaims), err
}