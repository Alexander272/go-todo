package domain

import (
	"time"
)

type TodoList struct {
	Id          string    `json:"id" bson:"_id,omitempty"`
	UserId      string    `json:"userId" bson:"userId,omitempty"`
	CategoryId  string    `json:"categoryId" bson:"categoryId,omitempty"`
	Title       string    `json:"title" bson:"title,omitempty"`
	Description string    `json:"description" bson:"description,omitempty"`
	CreatedAt   time.Time `json:"createdAt" bson:"createdAt,omitempty"`
	// Tags        []string  `json:"tags" bson:"tags"`
}

// todo можно потом добавить Files File

type CreateListDTO struct {
	Title       string `json:"title" binding:"required,min=3,max=128"`
	UserId      string `json:"userId"`
	Description string `json:"description"`
	CategoryId  string `json:"categoryId"`
}

func NewTodoList(dto CreateListDTO) TodoList {
	return TodoList{
		Title:       dto.Title,
		UserId:      dto.UserId,
		Description: dto.Description,
		CategoryId:  dto.CategoryId,
	}
}

type UpdateListDTO struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	CategoryId  string `json:"categoryId"`
}

func UpdateTodoList(dto UpdateListDTO) TodoList {
	return TodoList{
		Title:       dto.Title,
		Description: dto.Description,
		CategoryId:  dto.CategoryId,
	}
}
