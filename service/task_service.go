package service

import (
	"self_crud/models"
	"self_crud/repository"
)

type TaskService struct {
	repo repository.Task
}

func NewTaskService(repo repository.Task) *TaskService {
	return &TaskService{repo: repo}
}

func (t *TaskService) GetTask(id int) (*models.Task, error) {
	return t.repo.GetTask(id)
}

func (t *TaskService) GetTasks() ([]*models.Task, error) {
	return t.repo.GetTasks()
}

func (t *TaskService) CreateTask(task *models.Task) (int, error) {
	return t.repo.CreateTask(task)
}
func (t *TaskService) UpdateTask(id int) error {
	return t.repo.UpdateTask(id)
}
func (t *TaskService) DeleteTask(id int) error {
	return t.repo.DeleteTask(id)
}
