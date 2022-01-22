package service

import (
	"self_crud/models"
	"time"

	"gorm.io/gorm"
)

type TaskService struct {
}

func (t *TaskService) GetTask(id int) (*models.Task, error) {
	return &models.Task{
		Model: gorm.Model{
			ID:        uint(id),
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
			DeletedAt: gorm.DeletedAt{},
		},
		Title:  "Aboba",
		Status: "In progress",
	}, nil
}

func (t *TaskService) GetTasks() ([]*models.Task, error) {
	return []*models.Task{}, nil
}

func (t *TaskService) CreateTask(*models.Task) (int, error) {
	return 0, nil
}
func (t *TaskService) UpdateTask(id int) error {
	return nil
}
func (t *TaskService) DeleteTask(id int) error {
	return nil
}
