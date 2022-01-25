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

func (t *TaskRepository) GetTask(id int) (*models.Task, error) {
	res := new(models.Task)
	if err := t.db.Find(res, id).Error; err != nil{
		return nil, err
	}
	return res, nil

}
func (t *TaskRepository) GetTasks() ([]*models.Task, error) {
	var temp = make([]models.Task,0)
	if err := t.db.Find(&temp).Error; err != nil{
		return nil, err
	}
	res := make([]*models.Task, len(temp))
	for i := range temp{
		res[i] = &temp[i]
	}
	return res, nil
}
func (t *TaskRepository) CreateTask(task *models.Task) (int, error) {
	if err := t.db.Create(task).Error; err != nil{
		return 0, err
	}
	return int(task.ID), nil
}
func (t *TaskRepository) UpdateTask(id int) error {
	 return t.db.Model(&models.Task{}).Where("id = ?", id).Update("status", "done").Error
}

func (t *TaskRepository) DeleteTask(id int) error {
	return t.db.Delete(&models.Task{}, id).Error
}
