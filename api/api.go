package api

import (
	"github.com/gin-gonic/gin"
)

type VehicleApi interface {
	GetAllTask(ctx *gin.Context)
	GetTask(ctx *gin.Context)
	CreateTask(ctx *gin.Context)
	DeleteTask(ctx *gin.Context)
	UpdateTask(ctx *gin.Context)
}
