package service

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/Alexander272/go-todo/internal/domain"
	"github.com/Alexander272/go-todo/internal/repository"
	"github.com/Alexander272/go-todo/pkg/auth"
	"github.com/Alexander272/go-todo/pkg/hash"
)

type UserService struct {
	repo         repository.Users
	tokenManager auth.TokenManager
	hasher       hash.PasswordHasher
}

func NewUserService(repo repository.Users, tokenManager auth.TokenManager, hasher hash.PasswordHasher) *UserService {
	return &UserService{
		repo:         repo,
		tokenManager: tokenManager,
		hasher:       hasher,
	}
}

func (s *UserService) SignUp(ctx context.Context, dto domain.CreateUserDTO) (id string, err error) {
	candidate, _ := s.repo.GetByEmail(ctx, dto.Email)
	if (candidate != domain.User{}) {
		return id, domain.ErrUserAlreadyExists
	}
	passwordHash, err := s.hasher.HashPassword(dto.Password)
	if err != nil {
		return id, err
	}
	verificationCode, err := s.tokenManager.NewRefreshToken()
	if err != nil {
		return id, err
	}
	ttl, err := time.ParseDuration("6h")
	if err != nil {
		return id, err
	}

	user := domain.NewUser(dto)
	user.Password = passwordHash
	user.RegisteredAt = time.Now().Unix()
	user.LastVisitAt = time.Now().Unix()
	user.Verification = domain.Verification{
		Code:    verificationCode,
		Expires: time.Now().Add(ttl).Unix(),
	}

	id, err = s.repo.Create(ctx, user)
	if err != nil {
		if errors.Is(err, domain.ErrUserAlreadyExists) {
			return id, err
		}
		return id, err
	}

	return id, nil
}

func (s *UserService) GetAll(ctx context.Context) (users []domain.User, err error) {
	users, err = s.repo.GetAll(ctx)
	if err != nil {
		if errors.Is(err, domain.ErrUserNotFound) {
			return users, err
		}
		return users, fmt.Errorf("failed to get users. error: %w", err)
	}
	if len(users) == 0 {
		return users, domain.ErrUserNotFound
	}

	return users, nil
}

func (s *UserService) GetById(ctx context.Context, userId string) (u domain.User, err error) {
	u, err = s.repo.GetById(ctx, userId)
	if err != nil {
		if errors.Is(err, domain.ErrUserNotFound) {
			return u, err
		}
		return u, fmt.Errorf("failed to get user by id. error: %w", err)
	}

	return u, nil
}

func (s *UserService) Update(ctx context.Context, dto domain.UpdateUserDTO) error {
	updateUser := domain.UpdateUser(dto)
	err := s.repo.Update(ctx, updateUser)
	if err != nil {
		if errors.Is(err, domain.ErrUserNotFound) {
			return err
		}
		return fmt.Errorf("failed to update user. error: %w", err)
	}
	return nil
}

func (s *UserService) Remove(ctx context.Context, userId string) error {
	err := s.repo.Remove(ctx, userId)
	if err != nil {
		if errors.Is(err, domain.ErrUserNotFound) {
			return err
		}
		return fmt.Errorf("failed to remove user. error: %w", err)
	}
	return nil
}
