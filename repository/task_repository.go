package repository

import "self_crud/models"

type TaskRepository struct{}

func (t *TaskRepository) GetTask(id int) (*models.Task, error) {
	return nil, nil
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
