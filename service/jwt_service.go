package service

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"time"
)

type JWTService struct {
}

type claims struct {
	Id uint `json:"id"`
	jwt.StandardClaims
}

func (j *JWTService) GenerateToken(id uint) (string, error) {
	claims := &claims{
		id,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 12).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS384, claims)
	t, err := token.SignedString([]byte(signingKey))
	if err != nil {
		return "", err
	}
	return t, nil
}

func (j *JWTService) ParseToken(token string) (uint, error) {
	claims := &claims{}
	_, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("invalid signing method")
		}
		return []byte(signingKey), nil
	})
	if err != nil {
		return 0,  err
	}
	return claims.Id, nil
}
