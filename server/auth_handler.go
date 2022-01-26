package server

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"net/http"
	"self_crud/models"
	"self_crud/service"
	"strings"
)

func (h *Handler) signIn(c *gin.Context)  {
	var user models.User
	err := c.ShouldBindWith(&user, binding.JSON)
	if err != nil{
		sendErrorResponse(c, http.StatusBadRequest, "Invalid body " + err.Error())
		return
	}
	if err = validation(&user); err != nil{
		sendErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	token, err := h.service.GenerateToken(user.Username, user.Password)
	if err != nil{
		if err == service.NoSuchUser{
			sendErrorResponse(c, http.StatusNotFound, err.Error())
		}else{
			sendErrorResponse(c, http.StatusInternalServerError, err.Error())
		}
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"token":token,
	})

}


func (h *Handler) signUp(c *gin.Context)  {
	var user models.User
	err := c.ShouldBindWith(&user, binding.JSON)
	if err != nil{
		sendErrorResponse(c, http.StatusBadRequest, "Invalid body " + err.Error())
		return
	}
	if err = validation(&user); err != nil{
		sendErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	token, err := h.service.SignUp(user.Username, user.Password)
	if err != nil{
		if strings.Contains(err.Error(), "Duplicate entry"){
			sendErrorResponse(c, http.StatusConflict, "This username is already taken.")
		}else{
			sendErrorResponse(c, http.StatusInternalServerError, err.Error())
		}
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"token":token,
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