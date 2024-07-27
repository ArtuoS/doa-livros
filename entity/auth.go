package entity

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

var secretKey = "KLw1agcGnuZbZvtRjG"

type Auth struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Claims struct {
	UserId int64 `json:"user_id"`
	jwt.RegisteredClaims
}

func GenerateJWT(userId int64) (string, error) {
	mySigningKey := []byte(secretKey)
	claims := Claims{
		UserId: userId,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24)),
			Issuer:    "doa-livros",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(mySigningKey)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func ParseJWT(tokenString string) (*Claims, error) {
	mySigningKey := []byte(secretKey)

	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return mySigningKey, nil
	})

	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		return claims, nil
	} else {
		return nil, err
	}
}
