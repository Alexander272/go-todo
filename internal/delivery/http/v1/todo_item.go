package v1

import (
	"fmt"
	"net/http"
	"time"

	"github.com/Alexander272/go-todo/internal/service"
	"github.com/Alexander272/go-todo/pkg/logger"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (h *Handler) initItemRoutes(api *gin.RouterGroup) {
	api.GET("/:listId/todo", h.userIdentity, h.getItemsByListId)
	todo := api.Group("/todo", h.userIdentity)
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
// @Accept  json
// @Produce  json
// @Success 200 {array} domain.TodoItem
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /{listId}/todo [get]
func (h *Handler) getItemsByListId(c *gin.Context) {
	id := c.Param("listId")
	if id == "" {
		newErrorResponse(c, http.StatusBadRequest, "empty id param")
		return
	}
	listId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}
	ctxId, exists := c.Get(userIdCtx)
	if !exists {
		newErrorResponse(c, http.StatusForbidden, "access not allowed")
		return
	}
	userId, err := primitive.ObjectIDFromHex(fmt.Sprintf("%v", ctxId))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}

	lists, err := h.services.TodoItem.GetItemsByListId(c, userId, listId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, lists)
}

type CreateTodo struct {
	ListId      primitive.ObjectID `json:"listId" binding:"required"`
	Title       string             `json:"title" binding:"required,min=3,max=128"`
	Description string             `json:"description"`
	DeadlineAt  int64              `json:"deadlineAt"`
	Priority    int                `json:"priority"`
	Tags        []string           `json:"tags"`
}

// @Summary Create Item
// @Tags todo
// @Security ApiKeyAuth
// @Description создание задачи
// @ModuleID createItem
// @Accept  json
// @Produce  json
// @Param input body CreateTodo true "todo info"
// @Success 201 {object} idResponse
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /todo/ [post]
func (h *Handler) createItem(c *gin.Context) {
	ctxId, exists := c.Get(userIdCtx)
	if !exists {
		newErrorResponse(c, http.StatusForbidden, "access not allowed")
		return
	}
	userId, err := primitive.ObjectIDFromHex(fmt.Sprintf("%v", ctxId))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}

	var input CreateTodo
	if err := c.BindJSON(&input); err != nil {
		logger.Debug(err.Error())
		newErrorResponse(c, http.StatusBadRequest, "invalid input body")
		return
	}

	id, err := h.services.TodoItem.CreateItem(c, service.CreateTodoItem{
		UserId:      userId,
		ListId:      input.ListId,
		Title:       input.Title,
		Description: input.Description,
		DeadlineAt:  input.DeadlineAt,
		Priority:    input.Priority,
		Tags:        input.Tags,
	})
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusCreated, idResponse{Status: "Created", Id: id})
}

// @Summary Get All Items
// @Tags todo
// @Security ApiKeyAuth
// @Description получение всех задач пользователя
// @ModuleID getAllItems
// @Accept  json
// @Produce  json
// @Success 200 {array} domain.TodoItem
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /todo/all [get]
func (h *Handler) getAllItems(c *gin.Context) {
	ctxId, exists := c.Get(userIdCtx)
	if !exists {
		newErrorResponse(c, http.StatusForbidden, "access not allowed")
		return
	}
	userId, err := primitive.ObjectIDFromHex(fmt.Sprintf("%v", ctxId))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}

	items, err := h.services.TodoItem.GetItemsByUserId(c, userId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, items)
}

// @Summary Get Item By Id
// @Tags todo
// @Security ApiKeyAuth
// @Description получение одной задачи
// @ModuleID getItemById
// @Accept  json
// @Produce  json
// @Param id path string true "todo id"
// @Success 200 {object} domain.TodoItem
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /todo/{id} [get]
func (h *Handler) getItemById(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		newErrorResponse(c, http.StatusBadRequest, "empty id param")
		return
	}
	itemId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}

	item, err := h.services.TodoItem.GetItemsById(c, itemId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, item)
}

type UpdateTodo struct {
	ListId      primitive.ObjectID `json:"listId" binding:"required"`
	Title       string             `json:"title"`
	Description string             `json:"description"`
	DeadlineAt  time.Time          `json:"deadlineAt"`
	Priority    int                `json:"priority"`
	Done        bool               `json:"done"`
	Tags        []string           `json:"tags"`
}

// @Summary Update Item
// @Tags todo
// @Security ApiKeyAuth
// @Description обновление задачи
// @ModuleID updateItem
// @Accept  json
// @Produce  json
// @Param id path string true "todo id"
// @Param input body UpdateTodo true "todo info"
// @Success 200 {object} statusResponse
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /todo/{id} [patch]
func (h *Handler) updateItem(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		newErrorResponse(c, http.StatusBadRequest, "empty id param")
		return
	}
	itemId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}

	var input UpdateTodo
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid input body")
		return
	}

	if err := h.services.TodoItem.UpdateItem(c, service.UpdateTodoItem{
		Id:          itemId,
		ListId:      input.ListId,
		Title:       input.Title,
		Description: input.Description,
		DeadlineAt:  input.DeadlineAt,
		Priority:    input.Priority,
		Done:        input.Done,
		Tags:        input.Tags,
	}); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, statusResponse{"Updated"})
}

// @Summary Remove Item
// @Tags todo
// @Security ApiKeyAuth
// @Description удаление задачи
// @ModuleID removeItem
// @Accept  json
// @Produce  json
// @Param id path string true "todo id"
// @Success 200 {object} statusResponse
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /todo/{id} [delete]
func (h *Handler) removeItem(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		newErrorResponse(c, http.StatusBadRequest, "empty id param")
		return
	}
	itemId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}

	if err := h.services.TodoItem.RemoveItem(c, itemId); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, statusResponse{"Deleted"})
}
