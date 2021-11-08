package v1

import (
	"github.com/Alexander272/go-todo/pkg/logger"
	"github.com/gin-gonic/gin"
)

type errorResponse struct {
	Message string `json:"message"`
}

type statusResponse struct {
	Status string `json:"status"`
}

type idResponse struct {
	Status string      `json:"status"`
	Id     interface{} `json:"id"`
}

func newErrorResponse(c *gin.Context, statusCode int, message string) {
	logger.Errorf("Url: %s | ClientIp: %s | ErrorResponse: %s", c.Request.URL, c.ClientIP(), message)
	c.AbortWithStatusJSON(statusCode, errorResponse{message})
}
