package utils

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

const SECRETKEY = "supersecret"

func GenerateToken(email string, userId int64) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email":  email,
		"userId": userId,
		"exp":    time.Now().Add(time.Hour * 2).Unix(),
	})

	return token.SignedString([]byte(SECRETKEY))
}

func VerifyToken(token string) error {
	parsedToken, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		_, ok := t.Method.(*jwt.SigningMethodHMAC)

		if !ok {
			return nil, errors.New("Unexpected signin method")
		}

		return SECRETKEY, nil
	})

	if !parsedToken.Valid {
		return errors.New("Inavlid token")
	}

	// claims, ok := parsedToken.Claims.(jwt.MapClaims)
	// if !ok {
	// 	return errors.New("Imvalid token claims")
	// }

	// email, _ := claims["email"].(string)
	// userId, _ := claims["userId"].(int64)

	return nil
}
