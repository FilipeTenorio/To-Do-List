package service

import "todolist/models"

type VehicleService interface {
	RegisterVehicle(models.Vehicle) error
	UpdateVehicle(models.Vehicle) error
	GetAllVehicle() ([]models.Vehicle, error)
	DeleteVehicle(uint) error
}
