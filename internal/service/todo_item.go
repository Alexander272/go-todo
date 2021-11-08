package service

import (
	"context"
	"errors"
	"time"

	"github.com/Alexander272/go-todo/internal/domain"
	"github.com/Alexander272/go-todo/internal/repository"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type TodoItemService struct {
	repo repository.TodoItem
}

func NewTodoItemService(repo repository.TodoItem) *TodoItemService {
	return &TodoItemService{
		repo: repo,
	}
}

type CreateTodoItem struct {
	UserId      primitive.ObjectID
	ListId      primitive.ObjectID
	Title       string
	Description string
	DeadlineAt  int64
	Priority    int
	Tags        []string
}

func (s *TodoItemService) CreateItem(ctx context.Context, input CreateTodoItem) (interface{}, error) {
	candidate, err := s.repo.GetByTitle(ctx, input.UserId, input.Title)
	if err != nil {
		if !errors.Is(err, domain.ErrItemNotFound) {
			return nil, err
		}
	}
	if candidate != nil {
		return nil, domain.ErrItemAlreadyExists
	}

	item := domain.TodoItem{
		ListId:      input.ListId,
		UserId:      input.UserId,
		Title:       input.Title,
		Description: input.Description,
		CreatedAt:   time.Now().Unix(),
		DeadlineAt:  input.DeadlineAt,
		Done:        false,
		Priority:    input.Priority,
		Tags:        input.Tags,
	}
	return s.repo.Create(ctx, item)
}

func (s *TodoItemService) GetItemsByListId(ctx context.Context, userId, listId primitive.ObjectID) ([]domain.TodoItem, error) {
	return s.repo.GetByListId(ctx, userId, listId)
}

func (s *TodoItemService) GetItemsByUserId(ctx context.Context, userId primitive.ObjectID) ([]domain.TodoItem, error) {
	return s.repo.GetByUserId(ctx, userId)
}

func (s *TodoItemService) GetItemsById(ctx context.Context, itemId primitive.ObjectID) (*domain.TodoItem, error) {
	return s.repo.GetById(ctx, itemId)
}

type UpdateTodoItem struct {
	Id          primitive.ObjectID
	ListId      primitive.ObjectID
	Title       string
	Description string
	DeadlineAt  time.Time
	Priority    int
	Done        bool
	Tags        []string
}

func (s *TodoItemService) UpdateItem(ctx context.Context, input UpdateTodoItem) error {
	item := domain.TodoItem{
		Id:          input.Id,
		ListId:      input.ListId,
		Title:       input.Title,
		Description: input.Description,
		DeadlineAt:  input.DeadlineAt.Unix(),
		Priority:    input.Priority,
		Done:        input.Done,
		Tags:        input.Tags,
	}
	return s.repo.Update(ctx, item)
}

func (s *TodoItemService) RemoveItem(ctx context.Context, itemId primitive.ObjectID) error {
	return s.repo.Remove(ctx, itemId)
}
