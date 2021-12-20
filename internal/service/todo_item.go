package service

import (
	"context"
	"errors"
	"fmt"

	"github.com/Alexander272/go-todo/internal/domain"
	"github.com/Alexander272/go-todo/internal/repository"
)

type TodoItemService struct {
	repo repository.TodoItem
}

func NewTodoItemService(repo repository.TodoItem) *TodoItemService {
	return &TodoItemService{
		repo: repo,
	}
}

func (s *TodoItemService) Create(ctx context.Context, dto domain.CreateTodoDTO) (id string, err error) {
	candidate, err := s.repo.GetByTitle(ctx, dto.ListId, dto.Title)
	if err != nil {
		if !errors.Is(err, domain.ErrItemNotFound) {
			return id, fmt.Errorf("failed to get item by title. error: %w", err)
		}
	}
	if (candidate != domain.TodoItem{}) {
		return id, domain.ErrItemAlreadyExists
	}
	item := domain.NewTodo(dto)

	return s.repo.Create(ctx, item)
}

func (s *TodoItemService) GetByListId(ctx context.Context, listId string) (items []domain.TodoItem, err error) {
	items, err = s.repo.GetByListId(ctx, listId)
	if err != nil {
		if errors.Is(err, domain.ErrItemNotFound) {
			return items, err
		}
		return items, fmt.Errorf("failed to get items. error: %w", err)
	}
	if len(items) == 0 {
		return items, domain.ErrItemNotFound
	}

	return items, nil
}

func (s *TodoItemService) GetByUserId(ctx context.Context, userId string) (items []domain.TodoItem, err error) {
	items, err = s.repo.GetByUserId(ctx, userId)
	if err != nil {
		if errors.Is(err, domain.ErrItemNotFound) {
			return items, err
		}
		return items, fmt.Errorf("failed to get items. error: %w", err)
	}
	if len(items) == 0 {
		return items, domain.ErrItemNotFound
	}

	return items, nil
}

func (s *TodoItemService) GetById(ctx context.Context, itemId string) (item domain.TodoItem, err error) {
	item, err = s.repo.GetById(ctx, itemId)
	if err != nil {
		if errors.Is(err, domain.ErrItemNotFound) {
			return item, err
		}
		return item, fmt.Errorf("failed to get item by id. error: %w", err)
	}

	return item, nil
}

func (s *TodoItemService) Update(ctx context.Context, dto domain.UpdateTodoDTO) error {
	updateTodo := domain.UpdateTodo(dto)
	err := s.repo.Update(ctx, updateTodo)
	if err != nil {
		if errors.Is(err, domain.ErrItemNotFound) {
			return err
		}
		return fmt.Errorf("failed to update list. error: %w", err)
	}
	return nil
}

func (s *TodoItemService) Remove(ctx context.Context, itemId string) error {
	err := s.repo.Remove(ctx, itemId)
	if err != nil {
		if errors.Is(err, domain.ErrItemNotFound) {
			return err
		}
		return fmt.Errorf("failed to remove list. error: %w", err)
	}
	return nil
}
