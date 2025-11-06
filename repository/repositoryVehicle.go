package repository

import (
	"fmt"
	"os"
	"todolist/models"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type vehicleRepository struct {
	db *gorm.DB
}

func NewvehicleRepository() (VehicleRepository, error) {
	if err := godotenv.Load(); err != nil {
		fmt.Println("Erro ao carregar o .env")
	}

	connectString := os.Getenv("connection")
	db, err := gorm.Open(mysql.Open(connectString), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	db.AutoMigrate(&models.Vehicle{})

	fmt.Println("Banco conectado com GORM")
	return &vehicleRepository{db: db}, nil
}

func (r *vehicleRepository) CreateVehicle(Vehicle models.Vehicle) error {
	result := r.db.Create(&Vehicle)
	return result.Error
}

func (r *vehicleRepository) DeleteVehicle(id uint) error {
	result := r.db.Delete(id)
	return result.Error
}

func (r *vehicleRepository) GetVehicleByID(id uint) (*models.Vehicle, error) {
	var Vehicle models.Vehicle

	result := r.db.First(&Vehicle, id)

	if result.Error != nil {
		fmt.Println(result.Error)
		return nil, result.Error
	}

	return &Vehicle, nil
}

func (r *vehicleRepository) GetVehicles() ([]models.Vehicle, error) {
	var Vehicles []models.Vehicle

	result := r.db.Find(&Vehicles)
	return Vehicles, result.Error
}

func (r *vehicleRepository) UpdateVehicle(Vehicle models.Vehicle) error {
	result := r.db.Save(&Vehicle)
	return result.Error

}
