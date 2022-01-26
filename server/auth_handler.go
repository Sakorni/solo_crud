package server

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"self_crud/models"
	"strings"
)

func (h *Handler) signIn(c *gin.Context)  {
	var user models.User
	err := c.BindJSON(&user)
	if err != nil{
		sendErrorResponse(c, http.StatusBadRequest, "Invalid body " + err.Error())
		return
	}
	if err = validation(&user); err != nil{
		sendErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	id, err := h.service.SignIn(user.Username, user.Password)
	if err != nil{
		sendErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	if id == 0{
		sendErrorResponse(c, http.StatusNotFound, "No such user")
		return
	}
	c.JSON(http.StatusOK, map[string]uint{
		"id":id,
	})

}
func (h *Handler) signUp(c *gin.Context)  {
	var user models.User
	err := c.BindJSON(&user)
	if err != nil{
		sendErrorResponse(c, http.StatusBadRequest, "Invalid body " + err.Error())
		return
	}
	if err = validation(&user); err != nil{
		sendErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	id, err := h.service.SignUp(user.Username, user.Password)
	if err != nil{
		sendErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusCreated, map[string]uint{
		"id":id,
	})

}


func validation(user *models.User) error{
	if len(strings.TrimSpace(user.Username)) < 3{
		return fmt.Errorf("username is too short. It must consist more than 3 characters without leading and trailing spaces")
	}
	password := strings.TrimSpace(user.Password)
	if len(password) < 5{
		return fmt.Errorf("password is too short. It must contain at least 5 characters")
	}
	return nil
}