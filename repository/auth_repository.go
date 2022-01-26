package repository

import (
	"gorm.io/gorm"
	"self_crud/models"
)

type AuthRepository struct{
	db *gorm.DB
}

func NewAuthRepository(db *gorm.DB) *AuthRepository {
	return &AuthRepository{db: db}
}

func (a *AuthRepository) SignUp(username, hashedPassword string) (uint, error) {
	user := models.User{
		Username: username,
		Password: hashedPassword,
	}
	err := a.db.Create(&user).Error
	if err != nil{
		return 0, err
	}
	return user.ID, nil
}

func (a *AuthRepository) SignIn(username, hashedPassword string) (uint, error) {
	user := models.User{
		Username: username,
		Password: hashedPassword,
	}
	err := a.db.Where(&user).Find(&user).Error
	if err != nil{
		return 0, err
	}
	return user.ID, err
}

