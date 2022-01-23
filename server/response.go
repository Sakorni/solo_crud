package server

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type serverError struct {
	Message string
}

func sendErrorResponse(c *gin.Context, code int, message string) {
	logrus.Error(message)
	c.AbortWithStatusJSON(code, serverError{message})
}
