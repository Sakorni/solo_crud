package repository

import (
	"self_crud/models"

	"gorm.io/gorm"
)

type Task interface {
	GetTask(uid uint, id int) (*models.Task, error)
	GetTasks(uid uint) ([]*models.Task, error)
	CreateTask(task *models.Task) (int, error)
	UpdateTask(uid uint, id int) error
	DeleteTask(uid uint, id int) error
}

type Auth interface {
	SignIn(username, hashedPassword string) (uint, error)
	SignUp(username, hashedPassword string) (uint, error)
}

type Repository struct {
	Task
	Auth
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		NewTaskRepository(db),
		NewAuthRepository(db),
	}
}
