package server

import (
	"net/http"
	"self_crud/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h *Handler) getAllTasks(c *gin.Context) {
	var tasks []models.Task
	c.IndentedJSON(http.StatusOK, tasks)
}

func (h *Handler) getSingleTask(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, map[string]string{"error:": "Invalid value in param"})
		return
	}
	task, err := h.service.GetTask(id)
	c.IndentedJSON(http.StatusOK, task)

}

func (h *Handler) updateTask(c *gin.Context) {

}

func (h *Handler) createTask(c *gin.Context) {

}

func (h *Handler) deleteTask(c *gin.Context) {

}
