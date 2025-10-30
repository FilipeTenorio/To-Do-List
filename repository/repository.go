package repository

import (
	"todolist/models"
)

type TaskRepository interface {
	DeleteTask(id uint) error
	UpdateTask(task models.Task) error
	CreateTask(task models.Task) error
	GetTasks() ([]models.Task, error)
	GetTaskByID(id uint) (*models.Task, error)
}
