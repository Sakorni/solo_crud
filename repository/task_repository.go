package repository

import (
	"self_crud/models"
	"time"

	"gorm.io/gorm"
)

type TaskRepository struct {
	db *gorm.DB
}

func NewTaskRepository(db *gorm.DB) *TaskRepository {
	return &TaskRepository{db}
}

func (t *TaskRepository) GetTask(id int) (*models.Task, error) {
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
func (t *TaskRepository) GetTasks() ([]*models.Task, error) {
	return nil, nil
}
func (t *TaskRepository) CreateTask(*models.Task) (int, error) {
	return 0, nil
}
func (t *TaskRepository) UpdateTask(id int) error {
	return nil
}

func (t *TaskRepository) DeleteTask(id int) error {
	return nil
}
