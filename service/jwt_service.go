package service

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"time"
)

type JWTService struct {

}


type claims struct{
	id uint `json:"id"`
	jwt.StandardClaims
}


func (j *JWTService) GenerateToken(id uint) (string, error){
	claims := &claims{
		id,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 12).Unix(),
			IssuedAt: time.Now().Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS384, claims)
	t,err := token.SignedString(signingKey)
	if err != nil{
		return "", err
	}
	return t, nil
}


func (j *JWTService) ParseToken(token string) (uint, error){
	accessToken, err := jwt.ParseWithClaims(token, &claims{},func(token *jwt.Token)(interface{}, error){
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok{
			return nil, fmt.Errorf("invalid signing method")
		}
		return signingKey, nil
	})
	if err != nil{
		return 0, err
	}
	claims, ok := accessToken.Claims.(*claims)
	if !ok{
		return 0, fmt.Errorf("invalid type of claims")
	}
	return claims.id, nil
}