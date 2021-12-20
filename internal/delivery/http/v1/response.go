package v1

import (
	"github.com/Alexander272/go-todo/pkg/logger"
	"github.com/gin-gonic/gin"
)

type dataResponse struct {
	Data  interface{} `json:"data"`
	Count int64       `json:"count,omitempty"`
}

type idResponse struct {
	Id      string `json:"id,omitempty"`
	Message string `json:"message,omitempty"`
}

type errorResponse struct {
	Message string `json:"message"`
}

type statusResponse struct {
	Status string `json:"status"`
}

func newErrorResponse(c *gin.Context, statusCode int, message string) {
	logger.Errorf("Url: %s | ClientIp: %s | ErrorResponse: %s", c.Request.URL, c.ClientIP(), message)
	c.AbortWithStatusJSON(statusCode, errorResponse{message})
}
