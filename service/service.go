package service

import (
	"self_crud/models"
	"self_crud/repository"
)

//Service is an implementation of core funcions of the app.
type Task interface {
	GetTask(id int) (*models.Task, error)
	GetTasks() ([]*models.Task, error)
	CreateTask(*models.Task) (int, error)
	UpdateTask(id int) error
	DeleteTask(id int) error
}

type Auth interface {
	SignIn(username, password string) (uint, error)
	SignUp(username, password string) (uint, error)
}

type Service struct {
	Task
	Auth
}

func NewService(rep *repository.Repository) *Service {
	return &Service{
		NewTaskService(rep),
		NewAuthService(rep),
	}
}
