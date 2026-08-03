package auth

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var SecretKey = []byte("CHANGE_THIS_SECRET_IN_PRODUCTION")

func GenerateToken(userID string) (string, error) {

	token := jwt.NewWithClaims(
		jwt.SigningMethodHS256,
		jwt.MapClaims{
			"user_id": userID,
			"exp": time.Now().Add(24 * time.Hour).Unix(),
		},
	)

	return token.SignedString(SecretKey)
}
