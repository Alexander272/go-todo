package v1

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/Alexander272/go-todo/internal/domain"
	"github.com/gin-gonic/gin"
)

func (h *Handler) initCategoryRoutes(api *gin.RouterGroup) {
	category := api.Group("/categories", h.userIdentity)
	{
		category.GET("/", h.getAllCategories)
		category.POST("/", h.createCategory)
		category.PATCH("/:id", h.updateCategory)
		category.DELETE("/:id", h.removeCategory)
		category.GET("/lists", h.getCategoriesWithLists)
	}
}

// @Summary Create Category
// @Tags category
// @Security ApiKeyAuth
// @Description создание категории
// @ModuleID createCategory
// @Accept json
// @Produce json
// @Param input body domain.CreateCategoryDTO true "category info"
// @Success 201 {object} idResponse
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /categories/ [post]
func (h *Handler) createCategory(c *gin.Context) {
	userId, exists := c.Get(userIdCtx)
	if !exists {
		newErrorResponse(c, http.StatusForbidden, "access not allowed")
		return
	}

	var dto domain.CreateCategoryDTO
	if err := c.BindJSON(&dto); err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid input body")
		return
	}
	dto.UserId = userId.(string)

	id, err := h.services.Category.Create(c, dto)
	if err != nil {
		if errors.Is(err, domain.ErrCategoryAlreadyExists) {
			c.JSON(http.StatusBadRequest, errorResponse{err.Error()})
			return
		}
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.Header("Location", fmt.Sprintf("/api/v1/categories/%s", id))
	c.JSON(http.StatusCreated, idResponse{Id: id, Message: "Created"})
}

// @Summary Get All Categories
// @Security ApiKeyAuth
// @Tags category
// @Description получение всех категорий
// @ModuleID getAllCategories
// @Accept json
// @Produce json
// @Success 200 {object} dataResponse{data=[]domain.Category}
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /categories/ [get]
func (h *Handler) getAllCategories(c *gin.Context) {
	userId, exists := c.Get(userIdCtx)
	if !exists {
		newErrorResponse(c, http.StatusForbidden, "access not allowed")
		return
	}

	categories, err := h.services.Category.GetAll(c, userId.(string))
	if err != nil {
		if errors.Is(err, domain.ErrCategoryNotFound) {
			c.JSON(http.StatusNotFound, errorResponse{err.Error()})
			return
		}
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, dataResponse{Data: categories})
}

// @Summary Get Categories With Lists
// @Security ApiKeyAuth
// @Tags category
// @Description получение всех категорий и списков
// @ModuleID getCategoriesWithLists
// @Accept json
// @Produce json
// @Success 200 {object} dataResponse{data=[]domain.CategoryWithLists}
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /categories/lists/ [get]
func (h *Handler) getCategoriesWithLists(c *gin.Context) {
	userId, exists := c.Get(userIdCtx)
	if !exists {
		newErrorResponse(c, http.StatusForbidden, "access not allowed")
		return
	}

	categories, err := h.services.Category.GetWithLists(c, userId.(string))
	if err != nil {
		if errors.Is(err, domain.ErrCategoryNotFound) {
			c.JSON(http.StatusNotFound, errorResponse{err.Error()})
			return
		}
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, dataResponse{Data: categories})
}

// @Summary Update Category
// @Tags category
// @Security ApiKeyAuth
// @Description обновление данных категории
// @ModuleID updateCategory
// @Accept json
// @Produce json
// @Param id path string true "category id"
// @Param input body domain.UpdateCategoryDTO true "category info"
// @Success 200 {object} idResponse
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /categories/{id} [patch]
func (h *Handler) updateCategory(c *gin.Context) {
	categotyId := c.Param("id")
	if categotyId == "" {
		newErrorResponse(c, http.StatusBadRequest, "empty id param")
		return
	}

	var dto domain.UpdateCategoryDTO
	if err := c.BindJSON(&dto); err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid input body")
		return
	}
	dto.Id = categotyId

	if err := h.services.Category.Update(c, dto); err != nil {
		if errors.Is(err, domain.ErrCategoryNotFound) {
			c.JSON(http.StatusNotFound, errorResponse{err.Error()})
			return
		}
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, idResponse{Id: categotyId, Message: "Updated"})
}

// @Summary Remove Category
// @Tags category
// @Security ApiKeyAuth
// @Description удаление категории
// @ModuleID removeCategory
// @Accept json
// @Produce json
// @Param id path string true "category id"
// @Success 204 {object} idResponse
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /categories/{id} [delete]
func (h *Handler) removeCategory(c *gin.Context) {
	categotyId := c.Param("id")
	if categotyId == "" {
		newErrorResponse(c, http.StatusBadRequest, "empty id param")
		return
	}

	if err := h.services.Category.Remove(c, categotyId); err != nil {
		if errors.Is(err, domain.ErrCategoryNotFound) {
			c.JSON(http.StatusNotFound, errorResponse{err.Error()})
			return
		}
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusNoContent, idResponse{Message: "Deleted"})
}
