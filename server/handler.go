package server

import "github.com/gin-gonic/gin"

type Handler struct {
}

func (h *Handler) InitHandler() *gin.Engine {
	core := gin.New()

	taskGroup := core.Group("/tasks")
	{
		taskGroup.GET("/", h.getTasks)
		taskGroup.GET("/:id", h.getTasks)
		taskGroup.POST("/", h.createTask)
		taskGroup.PUT("/:id", h.updateTask)
		taskGroup.DELETE("/:id", h.deleteTask)
	}

	return core
}
