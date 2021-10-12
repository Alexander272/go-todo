package service

import (
	"context"
	"errors"
	"time"

	"github.com/Alexander272/go-todo/internal/domain"
	"github.com/Alexander272/go-todo/internal/repository"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type TodoListServise struct {
	repo     repository.TodoList
	repoItem repository.TodoItem
}

func NewTodoListService(repo repository.TodoList, repoItem repository.TodoItem) *TodoListServise {
	return &TodoListServise{
		repo:     repo,
		repoItem: repoItem,
	}
}

type CreateTodoList struct {
	UserId      primitive.ObjectID
	Title       string
	Description string
}

func (s *TodoListServise) CreateList(ctx context.Context, input CreateTodoList) error {
	candidate, err := s.repo.GetByTitle(ctx, input.UserId, input.Title)
	if !errors.Is(err, domain.ErrListNotFound) {
		return err
	}
	if candidate != nil {
		return domain.ErrListAlreadyExists
	}

	list := domain.TodoList{
		UserId:      input.UserId,
		Title:       input.Title,
		Description: input.Description,
		CreatedAt:   time.Now(),
	}
	return s.repo.Create(ctx, list)
}

func (s *TodoListServise) GetAllLists(ctx context.Context, userId primitive.ObjectID) ([]domain.TodoList, error) {
	return s.repo.GetAll(ctx, userId)
}

func (s *TodoListServise) GetListById(ctx context.Context, listId primitive.ObjectID) (*domain.TodoList, error) {
	return s.repo.GetById(ctx, listId)
}

type UpdateTodolist struct {
	Id          primitive.ObjectID
	Title       string
	Description string
}

func (s *TodoListServise) UpdateList(ctx context.Context, listId primitive.ObjectID, input UpdateTodolist) error {
	list := domain.TodoList{
		Id:          input.Id,
		Title:       input.Title,
		Description: input.Description,
	}
	return s.repo.Update(ctx, list)
}

func (s *TodoListServise) RemoveList(ctx context.Context, listId primitive.ObjectID) error {
	err := s.repoItem.RemoveByListId(ctx, listId)
	if err != nil {
		return err
	}
	return s.repo.Remove(ctx, listId)
}
