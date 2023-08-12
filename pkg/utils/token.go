package utils

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt"
)

type claim struct {
	id             uint
	email          string
	standardClaims jwt.StandardClaims
}

func CreateToken(duration time.Duration, id uint, email, secretKey string) (string, error) {
	// decodedPrivateKey, err := base64.StdEncoding.DecodeString(secretKey)
	// if err != nil {
	// 	fmt.Println("failed to decode")
	// 	return "", fmt.Errorf("failed to decode private key: %s", err.Error())
	// }

	// key, err := jwt.ParseRSAPrivateKeyFromPEM(decodedPrivateKey)
	// if err != nil {
	// 	fmt.Println("failed to parse decodedkey, \n" + err.Error())
	// 	return "", fmt.Errorf("failed to parse decoded private key: %s", err.Error())
	// }

	claims := &claim{
		id:    id,
		email: email,
		standardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(duration).Unix(),
		},
	}

	signedToken, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(secretKey))

	if err != nil {
		fmt.Println("failed to sign token")
		return "", fmt.Errorf("failed to create token, error: %s", err.Error())
	}

	return signedToken, nil
}

func VerifyToken(token, publicKey string) (interface{}, error) {
	// decodedPublicKey, err := base64.StdEncoding.DecodeString(publicKey)
	// if err != nil {
	// 	return "", fmt.Errorf("failed to decode public key")
	// }

	// key, err := jwt.ParseECPublicKeyFromPEM(decodedPublicKey)
	// if err != nil {
	// 	return "", fmt.Errorf("failed to parse public key: %s", err.Error())
	// }

	claims := claim{}
	parsedToken, err := jwt.ParseWithClaims(token, claims, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodRSA); !ok {
			return "", fmt.Errorf("unexpected method: %s", t.Header["alg"])
		}

		return []byte(publicKey), nil
	})

	if err != nil {
		return "", fmt.Errorf("validate: %s", err.Error())
	}

	if !parsedToken.Valid {
		return "", fmt.Errorf("validate: invalid token")
	}

	return claims.id, nil
}
