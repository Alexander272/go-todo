package v1

import (
	"fmt"
	"net/http"

	"github.com/Alexander272/go-todo/internal/service"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (h *Handler) initListRoutes(api *gin.RouterGroup) {
	list := api.Group("/list", h.userIdentity)
	{
		list.GET("/", h.getAllLists)
		list.POST("/", h.createList)
		list.GET("/:id", h.getListById)
		list.PUT("/:id", h.updateList)
		list.DELETE("/:id", h.removeList)
	}
}

// @Summary Get All Lists
// @Security ApiKeyAuth
// @Tags list
// @Description получение всех пользовательских списков
// @ModuleID getAllLists
// @Accept  json
// @Produce  json
// @Success 200 {array} domain.TodoList
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /list/ [get]
func (h *Handler) getAllLists(c *gin.Context) {
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

	lists, err := h.services.TodoList.GetAllLists(c, userId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, lists)
}

type CreateList struct {
	Title       string `json:"title" binding:"required,min=3,max=128"`
	Description string `json:"description"`
}

// @Summary Create List
// @Tags list
// @Security ApiKeyAuth
// @Description создание списка
// @ModuleID createList
// @Accept  json
// @Produce  json
// @Param input body CreateList true "list info"
// @Success 201 {object} statusResponse
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /list/ [post]
func (h *Handler) createList(c *gin.Context) {
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
	var inp CreateList
	if err := c.BindJSON(&inp); err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid input body")
		return
	}

	if err := h.services.TodoList.CreateList(c, service.CreateTodoList{
		UserId:      userId,
		Title:       inp.Title,
		Description: inp.Description,
	}); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusCreated, statusResponse{"Created"})
}

// @Summary Get List By Id
// @Tags list
// @Security ApiKeyAuth
// @Description получение данных списка
// @ModuleID getListById
// @Accept  json
// @Produce  json
// @Param id path string true "list id"
// @Success 200 {object} domain.TodoList
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /list/{id} [get]
func (h *Handler) getListById(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		newErrorResponse(c, http.StatusBadRequest, "empty id param")
		return
	}
	listId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}

	list, err := h.services.TodoList.GetListById(c, listId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, list)
}

type UpdateList struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}

// @Summary Update List
// @Tags list
// @Security ApiKeyAuth
// @Description обновление данных списка
// @ModuleID updateList
// @Accept  json
// @Produce  json
// @Param id path string true "list id"
// @Param input body UpdateList true "list info"
// @Success 200 {object} statusResponse
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /list/{id} [put]
func (h *Handler) updateList(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		newErrorResponse(c, http.StatusBadRequest, "empty id param")
		return
	}
	listId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}

	var input UpdateList
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid input body")
		return
	}

	if err := h.services.TodoList.UpdateList(c, listId, service.UpdateTodolist{
		Title:       input.Title,
		Description: input.Description,
	}); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, statusResponse{"Updated"})
}

// @Summary Remove List
// @Tags list
// @Security ApiKeyAuth
// @Description удаление списка
// @ModuleID removeList
// @Accept  json
// @Produce  json
// @Param id path string true "list id"
// @Success 200 {object} statusResponse
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /list/{id} [delete]
func (h *Handler) removeList(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		newErrorResponse(c, http.StatusBadRequest, "empty id param")
		return
	}
	listId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}

	if err := h.services.TodoList.RemoveList(c, listId); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, statusResponse{"Deleted"})
}
