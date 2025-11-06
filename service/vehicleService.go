package service

import (
	"todolist/models"
	"todolist/repository"
)

type vehicleService struct {
	db repository.VehicleRepository
}

func NewvehicleService(database repository.VehicleRepository) VehicleService {
	return &vehicleService{
		db: database,
	}
}

func (r *vehicleService) RegisterVehicle(Vehicle models.Vehicle) error {
	err := r.db.CreateVehicle(Vehicle)
	if err != nil {
		return err
	}
	return nil

}

func (r *vehicleService) DeleteVehicle(id uint) error {
	err := r.db.DeleteVehicle(id)
	if err != nil {
		panic(err)
	}
	return nil
}

func (r *vehicleService) GetAllVehicle() (Vehicles []models.Vehicle, err error) {
	Vehicles, err = r.db.GetVehicles()
	if err != nil {
		return nil, err
	}
	return Vehicles, nil

}

func (r *vehicleService) UpdateVehicle(Vehicle models.Vehicle) error {
	err := r.db.UpdateVehicle(Vehicle)
	if err != nil {
		return err
	}
	return nil
}
