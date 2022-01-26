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
	GenerateToken(username, password string) (string, error)
	SignUp(username, password string) (string, error)
}

type Service struct {
	Task
	Auth
}

func NewService(rep *repository.Repository, jwt *JWTService) *Service {
	return &Service{
		NewTaskService(rep),
		NewAuthService(rep, jwt),
	}
}
