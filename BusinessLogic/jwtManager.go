package BusinessLogic

import (
	"crypto/rand"
)

var secretKey, _ = generateRandomSecretKey()

func generateRandomSecretKey() ([]byte, error) {
	jwtKey := make([]byte, 32)
	if _, err := rand.Read(jwtKey); err != nil {
		return nil, err
	}

	return jwtKey, nil
}
