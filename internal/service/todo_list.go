package service

import (
	"context"
	"errors"
	"fmt"

	"github.com/Alexander272/go-todo/internal/domain"
	"github.com/Alexander272/go-todo/internal/repository"
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

func (s *TodoListServise) CreateList(ctx context.Context, dto domain.CreateListDTO) (id string, err error) {
	candidate, err := s.repo.GetByTitle(ctx, dto.UserId, dto.Title)
	if err != nil {
		if !errors.Is(err, domain.ErrListNotFound) {
			return id, fmt.Errorf("failed to get list by title. error: %w", err)
		}
	}
	if (candidate != domain.TodoList{}) {
		return id, domain.ErrListAlreadyExists
	}

	list := domain.NewTodoList(dto)

	return s.repo.Create(ctx, list)
}

func (s *TodoListServise) GetAllLists(ctx context.Context, userId string) (lists []domain.TodoList, err error) {
	lists, err = s.repo.GetAll(ctx, userId)
	if err != nil {
		if errors.Is(err, domain.ErrListNotFound) {
			return lists, err
		}
		return lists, fmt.Errorf("failed to get lists. error: %w", err)
	}
	if len(lists) == 0 {
		return lists, domain.ErrListNotFound
	}

	return lists, nil
}

func (s *TodoListServise) GetListById(ctx context.Context, listId string) (list domain.TodoList, err error) {
	list, err = s.repo.GetById(ctx, listId)
	if err != nil {
		if errors.Is(err, domain.ErrListNotFound) {
			return list, err
		}
		return list, fmt.Errorf("failed to get list by id. error: %w", err)
	}

	return list, nil
}

func (s *TodoListServise) UpdateList(ctx context.Context, listId string, dto domain.UpdateListDTO) error {
	updateList := domain.UpdateTodoList(dto)
	err := s.repo.Update(ctx, updateList)
	if err != nil {
		if errors.Is(err, domain.ErrListNotFound) {
			return err
		}
		return fmt.Errorf("failed to update list. error: %w", err)
	}
	return nil
}

func (s *TodoListServise) RemoveList(ctx context.Context, listId string) error {
	// err := s.repoItem.RemoveByListId(ctx, listId)
	// if err != nil {
	// 	return err
	// }
	// return s.repo.Remove(ctx, listId)
	err := s.repo.Remove(ctx, listId)
	if err != nil {
		if errors.Is(err, domain.ErrListNotFound) {
			return err
		}
		return fmt.Errorf("failed to remove list. error: %w", err)
	}
	return nil
}
