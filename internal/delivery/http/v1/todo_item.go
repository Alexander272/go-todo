package v1

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/Alexander272/go-todo/internal/domain"
	"github.com/gin-gonic/gin"
)

func (h *Handler) initItemRoutes(api *gin.RouterGroup) {
	api.GET("/:listId/todos", h.userIdentity, h.getItemsByListId)
	todo := api.Group("/todos", h.userIdentity)
	{
		todo.POST("/", h.createItem)
		todo.GET("/all", h.getAllItems)
		todo.GET("/:id", h.getItemById)
		todo.PATCH("/:id", h.updateItem)
		todo.DELETE("/:id", h.removeItem)
	}
}

// @Summary Get Items By List Id
// @Tags todo
// @Security ApiKeyAuth
// @Description получение задач конкретного списка
// @ModuleID getItemsByListId
// @Accept json
// @Produce json
// @Success 200 {object} dataResponse{data=[]domain.TodoItem}
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /{listId}/todos [get]
func (h *Handler) getItemsByListId(c *gin.Context) {
	listId := c.Param("listId")
	if listId == "" {
		newErrorResponse(c, http.StatusBadRequest, "empty id param")
		return
	}

	lists, err := h.services.TodoItem.GetByListId(c, listId)
	if err != nil {
		if errors.Is(err, domain.ErrItemNotFound) {
			c.JSON(http.StatusNotFound, errorResponse{err.Error()})
			return
		}
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, dataResponse{Data: lists})
}

// @Summary Create Item
// @Tags todo
// @Security ApiKeyAuth
// @Description создание задачи
// @ModuleID createItem
// @Accept json
// @Produce json
// @Param input body domain.CreateTodoDTO true "todo info"
// @Success 201 {object} idResponse
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /todos/ [post]
func (h *Handler) createItem(c *gin.Context) {
	var dto domain.CreateTodoDTO
	if err := c.BindJSON(&dto); err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid input body")
		return
	}

	id, err := h.services.TodoItem.Create(c, dto)
	if err != nil {
		if errors.Is(err, domain.ErrItemAlreadyExists) {
			c.JSON(http.StatusBadRequest, errorResponse{err.Error()})
			return
		}
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.Header("Location", fmt.Sprintf("/api/v1/todos/%s", id))
	c.JSON(http.StatusCreated, idResponse{Id: id, Message: "Created"})
}

// @Summary Get All Items
// @Tags todo
// @Security ApiKeyAuth
// @Description получение всех задач пользователя
// @ModuleID getAllItems
// @Accept json
// @Produce json
// @Success 200 {object} dataResponse{data=[]domain.TodoItem}
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /todos/all [get]
func (h *Handler) getAllItems(c *gin.Context) {
	userId, exists := c.Get(userIdCtx)
	if !exists {
		newErrorResponse(c, http.StatusForbidden, "access not allowed")
		return
	}

	items, err := h.services.TodoItem.GetByUserId(c, userId.(string))
	if err != nil {
		if errors.Is(err, domain.ErrItemNotFound) {
			c.JSON(http.StatusNotFound, errorResponse{err.Error()})
			return
		}
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, dataResponse{Data: items})
}

// @Summary Get Item By Id
// @Tags todo
// @Security ApiKeyAuth
// @Description получение одной задачи
// @ModuleID getItemById
// @Accept json
// @Produce json
// @Param id path string true "todo id"
// @Success 200 {object} dataResponse{data=domain.TodoItem}
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /todos/{id} [get]
func (h *Handler) getItemById(c *gin.Context) {
	itemId := c.Param("id")
	if itemId == "" {
		newErrorResponse(c, http.StatusBadRequest, "empty id param")
		return
	}

	item, err := h.services.TodoItem.GetById(c, itemId)
	if err != nil {
		if errors.Is(err, domain.ErrItemNotFound) {
			c.JSON(http.StatusNotFound, errorResponse{err.Error()})
			return
		}
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, dataResponse{Data: item})
}

// @Summary Update Item
// @Tags todo
// @Security ApiKeyAuth
// @Description обновление задачи
// @ModuleID updateItem
// @Accept json
// @Produce json
// @Param id path string true "todo id"
// @Param input body domain.UpdateTodoDTO true "todo info"
// @Success 200 {object} idResponse
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /todos/{id} [patch]
func (h *Handler) updateItem(c *gin.Context) {
	itemId := c.Param("id")
	if itemId == "" {
		newErrorResponse(c, http.StatusBadRequest, "empty id param")
		return
	}

	var dto domain.UpdateTodoDTO
	if err := c.BindJSON(&dto); err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid input body")
		return
	}
	dto.Id = itemId

	if err := h.services.TodoItem.Update(c, dto); err != nil {
		if errors.Is(err, domain.ErrItemNotFound) {
			c.JSON(http.StatusNotFound, errorResponse{err.Error()})
			return
		}
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, idResponse{Id: itemId, Message: "Updated"})
}

// @Summary Remove Item
// @Tags todo
// @Security ApiKeyAuth
// @Description удаление задачи
// @ModuleID removeItem
// @Accept json
// @Produce json
// @Param id path string true "todo id"
// @Success 204 {object} idResponse
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /todos/{id} [delete]
func (h *Handler) removeItem(c *gin.Context) {
	itemId := c.Param("id")
	if itemId == "" {
		newErrorResponse(c, http.StatusBadRequest, "empty id param")
		return
	}

	if err := h.services.TodoItem.Remove(c, itemId); err != nil {
		if errors.Is(err, domain.ErrItemNotFound) {
			c.JSON(http.StatusNotFound, errorResponse{err.Error()})
			return
		}
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusNoContent, idResponse{Message: "Deleted"})
}
