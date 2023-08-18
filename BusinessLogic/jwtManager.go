package BusinessLogic

import (
	"crypto/rand"
	"github.com/dgrijalva/jwt-go"
	"time"
)

var secretKey, _ = generateRandomSecretKey()

func generateRandomSecretKey() ([]byte, error) {
	jwtKey := make([]byte, 32)
	if _, err := rand.Read(jwtKey); err != nil {
		return nil, err
	}

	return jwtKey, nil
}

func GenerateJWTToken(username string) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["sub"] = username
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()

	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
