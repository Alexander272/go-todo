package service

import (
	"context"
	"errors"
	"fmt"

	"github.com/Alexander272/go-todo/internal/domain"
	"github.com/Alexander272/go-todo/internal/repository"
)

type CategoryService struct {
	repo     repository.Category
	repoList repository.TodoList
}

func NewCategoryService(repo repository.Category, repoList repository.TodoList) *CategoryService {
	return &CategoryService{
		repo:     repo,
		repoList: repoList,
	}
}

func (s *CategoryService) Create(ctx context.Context, dto domain.CreateCategoryDTO) (id string, err error) {
	candidate, err := s.repo.GetByTitle(ctx, dto.UserId, dto.Title)
	if err != nil {
		if !errors.Is(err, domain.ErrCategoryNotFound) {
			return id, fmt.Errorf("failed to get category by title. error: %w", err)
		}
	}
	if (candidate != domain.Category{}) {
		return id, domain.ErrItemAlreadyExists
	}
	c := domain.NewCategory(dto)

	id, err = s.repo.Create(ctx, c)
	if err != nil {
		return id, fmt.Errorf("failed to create category. error: %w", err)
	}

	return id, nil
}

func (s *CategoryService) GetAll(ctx context.Context, userId string) (categories []domain.Category, err error) {
	categories, err = s.repo.GetAll(ctx, userId)
	if err != nil {
		if errors.Is(err, domain.ErrCategoryNotFound) {
			return categories, err
		}
		return categories, fmt.Errorf("failed to get categories. error: %w", err)
	}
	if len(categories) == 0 {
		return categories, domain.ErrCategoryNotFound
	}

	return categories, nil
}

func (s *CategoryService) GetWithLists(ctx context.Context, userId string) (categories []domain.CategoryWithLists, err error) {
	categories, err = s.repo.GetWithLists(ctx, userId)
	if err != nil {
		if errors.Is(err, domain.ErrCategoryNotFound) {
			return categories, err
		}
		return categories, fmt.Errorf("failed to get categories. error: %w", err)
	}
	if len(categories) == 0 {
		return categories, domain.ErrCategoryNotFound
	}

	return categories, nil
}

func (s *CategoryService) Update(ctx context.Context, dto domain.UpdateCategoryDTO) error {
	updateCat := domain.UpdateCategory(dto)
	err := s.repo.Update(ctx, updateCat)
	if err != nil {
		if errors.Is(err, domain.ErrListNotFound) {
			return err
		}
		return fmt.Errorf("failed to update category. error: %w", err)
	}
	return nil
}

func (s *CategoryService) Remove(ctx context.Context, categoryId string) error {
	err := s.repoList.RemoveCatogoryId(ctx, categoryId)
	if err != nil {
		if !errors.Is(err, domain.ErrListNotFound) {
			return fmt.Errorf("failed to update lists by categoryid. error: %w", err)
		}
	}

	err = s.repo.Remove(ctx, categoryId)
	if err != nil {
		if errors.Is(err, domain.ErrCategoryNotFound) {
			return err
		}
		return fmt.Errorf("failed to remove category. error: %w", err)
	}
	return nil
}
