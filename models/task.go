package models

import (
	"gorm.io/gorm"
)

//There GORM provides me fields such as ID, CreatedAt, UpdatedAt, DeletedAt
type Task struct {
	gorm.Model
	Title  string `json:"title" db:"title"`
	Status string `json:"status" db:"status"`
	UserID uint   `json:"user_id"`
}
