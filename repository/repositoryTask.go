package repository

import (
	"fmt"
	"os"
	"todolist/models"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type taskRepository struct {
	db *gorm.DB
}

func NewTaskRepository() (TaskRepository, error) {
	if err := godotenv.Load(); err != nil {
		fmt.Println("Erro ao carregar o .env")
	}

	connectString := os.Getenv("connection")
	db, err := gorm.Open(mysql.Open(connectString), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	db.AutoMigrate(&models.Task{})

	fmt.Println("Banco conectado com GORM")
	return &taskRepository{db: db}, nil
}

func (r *taskRepository) CreateTask(task models.Task) error {
	result := r.db.Create(&task)
	return result.Error
}

func (r *taskRepository) DeleteTask(id uint) error {
	result := r.db.Delete(id)
	return result.Error
}

func (r *taskRepository) GetTaskByID(id uint) (*models.Task, error) {
	var task models.Task

	result := r.db.First(&task, id)

	if result.Error != nil {
		fmt.Println(result.Error)
		return nil, result.Error
	}

	return &task, nil
}

func (r *taskRepository) GetTasks() ([]*models.Task, error) {
	r.db.Get()
}
