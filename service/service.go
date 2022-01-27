package service

import (
	"self_crud/models"
	"self_crud/repository"
)

//Service is an implementation of core funcions of the app.
type Task interface {
	GetTask(uid uint, id int) (*models.Task, error)
	GetTasks(uid uint) ([]*models.Task, error)
	CreateTask(task *models.Task) (int, error)
	UpdateTask(uid uint, id int) error
	DeleteTask(uid uint, id int) error
}

type Auth interface {
	GenerateToken(username, password string) (string, error)
	SignUp(username, password string) (string, error)
	ParseToken(token string) (uint, error)
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
