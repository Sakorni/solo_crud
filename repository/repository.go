package repository

import "self_crud/models"

type Task interface {
	GetTask(id int) (*models.Task, error)
	GetTasks() ([]*models.Task, error)
	CreateTask(*models.Task) (int, error)
	UpdateTask(id int) error
	DeleteTask(id int) error
}

type Repository struct {
	Task
}

func NewRepository() *Repository {
	return &Repository{
		NewTaskRepository(),
	}
}
