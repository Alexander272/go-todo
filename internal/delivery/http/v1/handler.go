package v1

import (
	"github.com/Alexander272/go-todo/internal/service"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	services *service.Services
}

func NewHandler(services *service.Services) *Handler {
	return &Handler{
		services: services,
	}
}

func (h *Handler) Init(api *gin.RouterGroup) {
	v1 := api.Group("/v1")
	{
		h.initAuthRoutes(v1)
		h.initUserRoutes(v1)
		h.initCategoryRoutes(v1)
		h.initListRoutes(v1)
		h.initItemRoutes(v1)

		v1.GET("/", h.notImplemented)
	}
}

func (h *Handler) notImplemented(c *gin.Context) {}
