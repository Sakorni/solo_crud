package repository

import (
	"gorm.io/gorm"
	"self_crud/models"
)

type TaskRepository struct {
	db *gorm.DB
}

func NewTaskRepository(db *gorm.DB) *TaskRepository {
	return &TaskRepository{db}
}

func (t *TaskRepository) GetTask(uid uint, id int) (*models.Task, error) {
	res := new(models.Task)
	if err := t.db.Where(&models.Task{UserID: uid}).Find(res, id).Error; err != nil {
		return nil, err
	}
	return res, nil

}
func (t *TaskRepository) GetTasks(uid uint) ([]*models.Task, error) {
	var res []*models.Task
	if err := t.db.Where(&models.Task{UserID: uid}).Find(&res).Error; err != nil {
		return nil, err
	}
	return res, nil
}
func (t *TaskRepository) CreateTask(task *models.Task) (int, error) {
	if err := t.db.Create(task).Error; err != nil {
		return 0, err
	}
	return int(task.ID), nil
}
func (t *TaskRepository) UpdateTask(uid uint, id int) error {
	return t.db.Model(&models.Task{}).Where("user_id = ? and id = ?", uid, id).Update("status", "done").Error
}

func (t *TaskRepository) DeleteTask(uid uint, id int) error {
	return t.db.Where(&models.Task{UserID: uid}).Delete(&models.Task{}, id).Error
}
