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

func (t *TaskService) GetTask(uid uint, id int) (*models.Task, error) {
	return t.repo.GetTask(uid, id)
}

func (t *TaskService) GetTasks(uid uint) ([]*models.Task, error) {
	return t.repo.GetTasks(uid)
}

func (t *TaskService) CreateTask(task *models.Task) (int, error) {
	return t.repo.CreateTask(task)
}
func (t *TaskService) UpdateTask(uid uint, id int) error {
	return t.repo.UpdateTask(uid, id)
}
func (t *TaskService) DeleteTask(uid uint, id int) error {
	return t.repo.DeleteTask(uid, id)
}
