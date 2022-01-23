package server

import (
	"fmt"
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
		sendErrorResponse(c, http.StatusBadRequest, "Invalid value in parameter id")
		return
	}
	task, err := h.service.GetTask(id)
	if err != nil {
		sendErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, task)

}

func (h *Handler) updateTask(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		sendErrorResponse(c, http.StatusBadRequest, "Invalid value in parameter id")
		return
	}
	err = h.service.UpdateTask(id)
	if err != nil {
		sendErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.Writer.WriteHeader(http.StatusOK)
}

func (h *Handler) createTask(c *gin.Context) {
	var task models.Task
	err := c.BindJSON(&task)
	if err != nil {
		sendErrorResponse(c, http.StatusBadRequest, fmt.Sprintf("Invalid body. %s", err.Error()))
		return
	}
	id, err := h.service.CreateTask(&task)
	if err != nil {
		sendErrorResponse(c, http.StatusInternalServerError, err.Error())
	}
	c.JSON(http.StatusOK, map[string]int{"id": id})
}

func (h *Handler) deleteTask(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		sendErrorResponse(c, http.StatusBadRequest, "Invalid value in parameter id")
		return
	}
	err = h.service.DeleteTask(id)
	if err != nil {
		sendErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	c.Writer.WriteHeader(http.StatusOK)
}
