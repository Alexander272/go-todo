package v1

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/Alexander272/go-todo/internal/domain"
	"github.com/gin-gonic/gin"
)

func (h *Handler) initListRoutes(api *gin.RouterGroup) {
	list := api.Group("/lists", h.userIdentity)
	{
		list.GET("/", h.getAllLists)
		list.POST("/", h.createList)
		list.GET("/:id", h.getListById)
		list.PATCH("/:id", h.updateList)
		list.DELETE("/:id", h.removeList)
	}
}

// @Summary Get All Lists
// @Security ApiKeyAuth
// @Tags list
// @Description получение всех пользовательских списков
// @ModuleID getAllLists
// @Accept json
// @Produce json
// @Success 200 {object} dataResponse{data=[]domain.TodoList}
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /lists/ [get]
func (h *Handler) getAllLists(c *gin.Context) {
	userId, exists := c.Get(userIdCtx)
	if !exists {
		newErrorResponse(c, http.StatusForbidden, "access not allowed")
		return
	}

	lists, err := h.services.TodoList.GetAll(c, userId.(string))
	if err != nil {
		if errors.Is(err, domain.ErrListNotFound) {
			c.JSON(http.StatusNotFound, errorResponse{err.Error()})
			return
		}
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, dataResponse{Data: lists})
}

// @Summary Create List
// @Tags list
// @Security ApiKeyAuth
// @Description создание списка
// @ModuleID createList
// @Accept json
// @Produce json
// @Param input body domain.CreateListDTO true "list info"
// @Success 201 {object} idResponse
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /lists/ [post]
func (h *Handler) createList(c *gin.Context) {
	userId, exists := c.Get(userIdCtx)
	if !exists {
		newErrorResponse(c, http.StatusForbidden, "access not allowed")
		return
	}

	var dto domain.CreateListDTO
	if err := c.BindJSON(&dto); err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid input body")
		return
	}
	dto.UserId = userId.(string)

	id, err := h.services.TodoList.Create(c, dto)
	if err != nil {
		if errors.Is(err, domain.ErrListAlreadyExists) {
			c.JSON(http.StatusBadRequest, errorResponse{err.Error()})
			return
		}
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.Header("Location", fmt.Sprintf("/api/v1/lists/%s", id))
	c.JSON(http.StatusCreated, idResponse{Id: id, Message: "Created"})
}

// @Summary Get List By Id
// @Tags list
// @Security ApiKeyAuth
// @Description получение данных списка
// @ModuleID getListById
// @Accept json
// @Produce json
// @Param id path string true "list id"
// @Success 200 {object} dataResponse{data=domain.TodoList}
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /lists/{id} [get]
func (h *Handler) getListById(c *gin.Context) {
	listId := c.Param("id")
	if listId == "" {
		newErrorResponse(c, http.StatusBadRequest, "empty id param")
		return
	}

	list, err := h.services.TodoList.GetById(c, listId)
	if err != nil {
		if errors.Is(err, domain.ErrListNotFound) {
			c.JSON(http.StatusNotFound, errorResponse{err.Error()})
			return
		}
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, dataResponse{Data: list})
}

// @Summary Update List
// @Tags list
// @Security ApiKeyAuth
// @Description обновление данных списка
// @ModuleID updateList
// @Accept json
// @Produce json
// @Param id path string true "list id"
// @Param input body domain.UpdateListDTO true "list info"
// @Success 200 {object} idResponse
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /lists/{id} [patch]
func (h *Handler) updateList(c *gin.Context) {
	listId := c.Param("id")
	if listId == "" {
		newErrorResponse(c, http.StatusBadRequest, "empty id param")
		return
	}

	var dto domain.UpdateListDTO
	if err := c.BindJSON(&dto); err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid input body")
		return
	}
	dto.Id = listId

	if err := h.services.TodoList.Update(c, dto); err != nil {
		if errors.Is(err, domain.ErrListNotFound) {
			c.JSON(http.StatusNotFound, errorResponse{err.Error()})
			return
		}
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, idResponse{Id: listId, Message: "Updated"})
}

// @Summary Remove List
// @Tags list
// @Security ApiKeyAuth
// @Description удаление списка
// @ModuleID removeList
// @Accept json
// @Produce json
// @Param id path string true "list id"
// @Success 204 {object} idResponse
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /lists/{id} [delete]
func (h *Handler) removeList(c *gin.Context) {
	listId := c.Param("id")
	if listId == "" {
		newErrorResponse(c, http.StatusBadRequest, "empty id param")
		return
	}

	if err := h.services.TodoList.Remove(c, listId); err != nil {
		if errors.Is(err, domain.ErrListNotFound) {
			c.JSON(http.StatusNotFound, errorResponse{err.Error()})
			return
		}
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusNoContent, idResponse{Message: "Deleted"})
}
