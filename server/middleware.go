package server

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func (h *Handler) identifyUser(c *gin.Context) {
	header := c.GetHeader("Authorization")
	if header == "" {
		sendErrorResponse(c, http.StatusUnauthorized, "empty auth header")
		return
	}
	headerParts := strings.Split(header, " ")
	if len(headerParts) != 2 {
		sendErrorResponse(c, http.StatusUnauthorized, "invalid auth header")
		return
	}
	id, err := h.service.ParseToken(headerParts[1])
	if err != nil {
		sendErrorResponse(c, http.StatusUnauthorized, err.Error())
		return
	}
	c.Set(uidKey, id)
}
