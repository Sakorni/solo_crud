package server

import (
	"self_crud/service"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	service *service.Service
}

func NewHandler(s *service.Service) *Handler {
	return &Handler{
		s,
	}
}

func (h *Handler) InitHandler() *gin.Engine {
	core := gin.New()

	taskGroup := core.Group("/tasks")
	{
		taskGroup.GET("/", h.getAllTasks)
		taskGroup.GET("/:id", h.getSingleTask)
		taskGroup.POST("/", h.createTask)
		taskGroup.PUT("/:id", h.updateTask)
		taskGroup.DELETE("/:id", h.deleteTask)
	}

	return core
}
