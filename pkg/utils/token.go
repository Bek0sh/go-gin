package utils

import (
	"encoding/base64"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt"
)

func CreateToken(duration time.Duration, payload interface{}, email, secretKey string) (string, error) {
	decodedPrivateKey, err := base64.StdEncoding.DecodeString(secretKey)
	if err != nil {
		fmt.Println("failed to decode")
		return "", fmt.Errorf("failed to decode private key: %s", err.Error())
	}

	key, err := jwt.ParseRSAPrivateKeyFromPEM(decodedPrivateKey)
	if err != nil {
		fmt.Println("failed to parse decodedkey, \n" + err.Error())
		return "", fmt.Errorf("failed to parse decoded private key: %s", err.Error())
	}

	claims := make(jwt.MapClaims)
	claims["sub"] = payload
	claims["exp"] = time.Hour * 24

	signedToken, err := jwt.NewWithClaims(jwt.SigningMethodRS256, claims).SignedString(key)

	if err != nil {
		fmt.Println("failed to sign token")
		return "", fmt.Errorf("failed to create token, error: %s", err.Error())
	}

	return signedToken, nil
}

func VerifyToken(token, publicKey string) (interface{}, error) {
	decodedPublicKey, err := base64.StdEncoding.DecodeString(publicKey)
	if err != nil {
		return 0, fmt.Errorf("failed to decode public key")
	}

	key, err := jwt.ParseRSAPublicKeyFromPEM(decodedPublicKey)
	if err != nil {
		return 0, fmt.Errorf("failed to parse public key: %s", err.Error())
	}

	parsedToken, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, fmt.Errorf("unexpected method: %s", t.Header["alg"])
		}
		return key, nil
	})

	if err != nil {
		fmt.Println("error 4 \n " + err.Error())
		return 0, fmt.Errorf("validate: %s", err.Error())
	}

	claims, ok := parsedToken.Claims.(jwt.MapClaims)

	if !parsedToken.Valid || !ok {
		return 0, fmt.Errorf("validate: invalid token")
	}

	return claims["sub"], nil
}
