package server

import (
	"self_crud/service"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	service *service.Service
}

const uidKey = "user_id"

func NewHandler(s *service.Service) *Handler {
	return &Handler{
		s,
	}
}

func (h *Handler) InitHandler() *gin.Engine {
	core := gin.New()

	taskGroup := core.Group("/tasks", h.identifyUser)
	{
		taskGroup.GET("/", h.getAllTasks)
		taskGroup.GET("/:id", h.getSingleTask)
		taskGroup.POST("/", h.createTask)
		taskGroup.PUT("/:id", h.updateTask)
		taskGroup.DELETE("/:id", h.deleteTask)
	}
	authGroup := core.Group("/auth")
	{
		authGroup.POST("/sign_in", h.signIn)
		authGroup.POST("/sign_up", h.signUp)
	}
	return core
}
