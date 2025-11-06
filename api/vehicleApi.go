package api

import (
	"todolist/models"
	"todolist/service"

	"github.com/gin-gonic/gin"
)

type vehicleApi struct {

	service service.TaskService
}

func newVehicleApi(services service.TaskService) VehicleApi {
	return &vehicleApi{
		service: services,
	}
}

func (i *taskApi) GetAllVehicle(ctx *gin.Context) {

}
func (i *taskApi) GetVehicle(ctx *gin.Context) {

}
func (i *taskApi)CreateVehicle(ctx *gin.Context) {

}
func (i *taskApi) DeleteVehicle(ctx *gin.Context) {

}
func (i *taskApi) UpdateVehicle(ctx *gin.Context) {

}

func checkTask(task models.Task){
	array = ["Carro","Moto","Caminhão","Van","Bicicleta"]
	for i := 0; i < count; i++ {
		
	}
}