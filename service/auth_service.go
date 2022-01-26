package service

import (
	"crypto/sha1"
	"fmt"
	"self_crud/repository"
)

const salt = "jsn12mslz.apdcks"

type AuthService struct{
	rep repository.Auth
}

func NewAuthService(rep repository.Auth) *AuthService {
	return &AuthService{
		rep,
	}
}

func (a *AuthService) SignIn(username, password string) (uint, error) {
	return a.rep.SignIn(username, hashPassword(password))
}

func (a *AuthService) SignUp(username, password string) (uint, error) {
	return a.rep.SignUp(username, hashPassword(password))
}

func hashPassword(password string) string{
	h := sha1.New()
	h.Write([]byte(salt))
	return fmt.Sprintf("%x",h.Sum([]byte(password)))
}