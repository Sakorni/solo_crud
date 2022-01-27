package server

import (
	"fmt"
	"net/http"
	"self_crud/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

func getUid(c *gin.Context) uint{
	uid, exists := c.Get(uidKey)
	if !exists{
		sendErrorResponse(c, http.StatusInternalServerError, "user id not found")
		c.Abort()
	}
	return uid.(uint)
}

func (h *Handler) getAllTasks(c *gin.Context) {
	uid := getUid(c)
	tasks, err := h.service.GetTasks(uid)
	if err != nil {
		sendErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data":tasks,
	})
}

func (h *Handler) getSingleTask(c *gin.Context) {
	uid := getUid(c)
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		sendErrorResponse(c, http.StatusBadRequest, "Invalid value in parameter id")
		return
	}
	task, err := h.service.GetTask(uid,id)
	if err != nil {
		sendErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, task)

}

func (h *Handler) updateTask(c *gin.Context) {
	uid := getUid(c)
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		sendErrorResponse(c, http.StatusBadRequest, "Invalid value in parameter id")
		return
	}
	err = h.service.UpdateTask(uid, id)
	if err != nil {
		sendErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.Writer.WriteHeader(http.StatusOK)
}

func (h *Handler) createTask(c *gin.Context) {
	uid := getUid(c)
	var task models.Task
	err := c.BindJSON(&task)
	if err != nil {
		sendErrorResponse(c, http.StatusBadRequest, fmt.Sprintf("Invalid body. %s", err.Error()))
		return
	}
	task.UserID = uid
	id, err := h.service.CreateTask(&task)
	if err != nil {
		sendErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusCreated, map[string]int{"id": id})
}

func (h *Handler) deleteTask(c *gin.Context) {
	uid := getUid(c)
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		sendErrorResponse(c, http.StatusBadRequest, "Invalid value in parameter id")
		return
	}
	err = h.service.DeleteTask(uid, id)
	if err != nil {
		sendErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	c.Writer.WriteHeader(http.StatusOK)
}
