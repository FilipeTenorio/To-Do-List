package repository

import "todolist/models"

type VehicleRepository interface {
	DeleteVehicle(id uint) error
	UpdateVehicle(vehicle models.Vehicle) error
	CreateVehicle(vehicle models.Vehicle) error
	GetVehicles() ([]models.Vehicle, error)
	GetVehicleByID(id uint) (*models.Vehicle, error)
}
