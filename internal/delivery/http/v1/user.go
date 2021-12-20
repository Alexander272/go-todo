package v1

import (
	"errors"
	"net/http"

	"github.com/Alexander272/go-todo/internal/domain"
	"github.com/gin-gonic/gin"
)

func (h *Handler) initUserRoutes(api *gin.RouterGroup) {
	user := api.Group("/users", h.userIdentity)
	{
		user.GET("/all", h.getAllUsers)
		user.GET("/:id", h.getUserById)
		user.PUT("/:id", h.updateUser)
		user.DELETE("/:id", h.removeUser)
	}
}

func (h *Handler) getAllUsers(c *gin.Context) {
	users, err := h.services.User.GetAll(c)
	if err != nil {
		if errors.Is(err, domain.ErrUserNotFound) {
			c.JSON(http.StatusNotFound, errorResponse{err.Error()})
			return
		}
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, users)
}

// @Summary Get User By Id
// @Security ApiKeyAuth
// @Tags user
// @Description получение данных пользователя
// @ModuleID getUserById
// @Accept json
// @Produce json
// @Param id path string true "user id"
// @Success 200 {object} dataResponse{data=domain.User}
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /users/{id} [get]
func (h *Handler) getUserById(c *gin.Context) {
	userId := c.Param("id")
	if userId == "" {
		newErrorResponse(c, http.StatusBadRequest, "empty id param")
		return
	}

	user, err := h.services.User.GetById(c, userId)
	if err != nil {
		if errors.Is(err, domain.ErrUserNotFound) {
			c.JSON(http.StatusNotFound, errorResponse{err.Error()})
			return
		}
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, dataResponse{Data: user})
}

// @Summary Update User
// @Security ApiKeyAuth
// @Tags user
// @Description обновление данных пользователя по его id
// @ModuleID updateUser
// @Accept json
// @Produce json
// @Param id path string true "user id"
// @Param input body domain.UpdateUserDTO true "user info"
// @Success 200 {object} idResponse
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /users/{id} [put]
func (h *Handler) updateUser(c *gin.Context) {
	userId := c.Param("id")
	if userId == "" {
		newErrorResponse(c, http.StatusBadRequest, "empty id param")
		return
	}

	var dto domain.UpdateUserDTO
	if err := c.Bind(&dto); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	dto.UserId = userId

	err := h.services.User.Update(c, dto)
	if err != nil {
		if errors.Is(err, domain.ErrUserNotFound) {
			c.JSON(http.StatusNotFound, errorResponse{err.Error()})
			return
		}
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, idResponse{Id: userId, Message: "Updated"})
}

// @Summary Remove User By Id
// @Security ApiKeyAuth
// @Tags user
// @Description удаление пользователя и всех его данных
// @ModuleID removeUserById
// @Accept  json
// @Produce  json
// @Param id path string true "user id"
// @Success 204 {object} idResponse
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /users/{id} [delete]
func (h *Handler) removeUser(c *gin.Context) {
	userId := c.Param("id")
	if userId == "" {
		newErrorResponse(c, http.StatusBadRequest, "empty id param")
		return
	}

	if err := h.services.User.Remove(c, userId); err != nil {
		if errors.Is(err, domain.ErrUserNotFound) {
			c.JSON(http.StatusNotFound, errorResponse{err.Error()})
			return
		}
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusNoContent, idResponse{Message: "Removed"})
}
