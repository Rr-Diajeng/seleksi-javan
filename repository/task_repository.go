package repository

import (
	"fmt"
	"seleksi-javan/model"

	"gorm.io/gorm"
)

type (
	TaskRepository interface {
		CreateTask(task model.Task) error
		FindTaskByID(taskId uint) (model.Task, error)
		GetAllTask() ([]model.Task, error)
		UpdateTask(taskId uint, updateTask model.Task) error
		DeleteTask(taskId uint) error
	}

	taskRepository struct {
		db *gorm.DB
	}
)

func NewTaskRepository(db *gorm.DB) TaskRepository {
	return &taskRepository{
		db: db,
	}
}

func (tr *taskRepository) CreateTask(task model.Task) error {
	err := tr.db.Create(&task).Error

	return err
}

func (tr *taskRepository) FindTaskByID(taskId uint) (model.Task, error) {
	task := model.Task{}
	err := tr.db.Where("id = ?", taskId).First(&task).Error

	return task, err
}

func (tr *taskRepository) GetAllTask() ([]model.Task, error) {
	var tasks []model.Task
	err := tr.db.Find(&tasks).Error

	return tasks, err
}

func (tr *taskRepository) UpdateTask(taskId uint, updateTask model.Task) error {
	tx := tr.db.Model(&model.Task{}).Where("id = ?", taskId).Updates(updateTask)

	if tx.Error != nil {
		return tx.Error
	}

	if tx.RowsAffected == 0 {
		return fmt.Errorf("task with id %d not found", taskId)
	}

	return nil
}

func (tr *taskRepository) DeleteTask(taskId uint) error {
	err := tr.db.Delete(&model.Task{}, taskId).Error

	return err
}
