package api

import (
	"github.com/gin-gonic/gin"
)

type TaskApi interface {
	CreateTask(ctx *gin.Context)
	DeleteTask(ctx *gin.Context)
	UpdateTask(ctx *gin.Context)
}
