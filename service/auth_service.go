package service

import (
	"crypto/sha1"
	"fmt"
	"self_crud/repository"
)

var NoSuchUser = fmt.Errorf("no such user")

const (
	salt       = "jsn12mslz.apdcks"
	signingKey = "kaso2lw223xla;wpxSDKzlspSNwPX"
)

type AuthService struct {
	rep        repository.Auth
	jwtService *JWTService
}

func NewAuthService(rep repository.Auth, jwt *JWTService) *AuthService {
	return &AuthService{
		rep,
		jwt,
	}
}

func (a *AuthService) GenerateToken(username, password string) (string, error) {
	id, err := a.rep.SignIn(username, hashPassword(password))
	if err != nil {
		return "", err
	}
	if id == 0 {
		return "", NoSuchUser
	}
	return a.jwtService.GenerateToken(id)
}

func (a *AuthService) SignUp(username, password string) (string, error) {
	id, err := a.rep.SignUp(username, hashPassword(password))
	if err != nil {
		return "", err
	}
	return a.jwtService.GenerateToken(id)
}

func (a *AuthService) ParseToken(token string) (uint, error) {
	return a.jwtService.ParseToken(token)
}

func hashPassword(password string) string {
	h := sha1.New()
	h.Write([]byte(salt))
	return fmt.Sprintf("%x", h.Sum([]byte(password)))
}
