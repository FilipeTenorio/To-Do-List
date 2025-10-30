package service

import "todolist/models"

type TaskService interface {
	RegisterTask(models.Task) error
	UpdateTask(models.Task) error
	GetAllTask() ([]models.Task, error)
	DeleteTask(string) error
}
