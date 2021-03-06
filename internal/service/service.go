package service

import (
	"context"
	"net/http"
	"time"

	"github.com/Alexander272/go-todo/internal/domain"
	"github.com/Alexander272/go-todo/internal/repository"
	"github.com/Alexander272/go-todo/pkg/auth"
	"github.com/Alexander272/go-todo/pkg/hash"
)

const CookieName = "session"

type SignInInput struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6,max=64"`
}

// type Token struct {
// 	AccessToken  string
// 	RefreshToken string
// }

type Session interface {
	SignIn(ctx context.Context, input SignInInput, ua, ip string) (*http.Cookie, *domain.Token, error)
	SingOut(ctx context.Context, token string) (*http.Cookie, error)
	Refresh(ctx context.Context, token, ua, ip string) (*domain.Token, *http.Cookie, error)
	TokenParse(token string) (userId string, role string, err error)
}

type User interface {
	SignUp(ctx context.Context, dto domain.CreateUserDTO) (string, error)
	GetAll(ctx context.Context) ([]domain.User, error)
	GetById(ctx context.Context, userId string) (domain.User, error)
	Update(ctx context.Context, dto domain.UpdateUserDTO) error
	Remove(ctx context.Context, userId string) error
}

type Category interface {
	Create(ctx context.Context, dto domain.CreateCategoryDTO) (string, error)
	GetAll(ctx context.Context, userId string) ([]domain.Category, error)
	GetWithLists(ctx context.Context, userId string) ([]domain.CategoryWithLists, error)
	Update(ctx context.Context, dto domain.UpdateCategoryDTO) error
	Remove(ctx context.Context, categoryId string) error
}

type TodoList interface {
	Create(ctx context.Context, dto domain.CreateListDTO) (string, error)
	GetAll(ctx context.Context, userId string) ([]domain.TodoList, error)
	GetById(ctx context.Context, listId string) (domain.TodoList, error)
	Update(ctx context.Context, dto domain.UpdateListDTO) error
	Remove(ctx context.Context, listId string) error
}

type TodoItem interface {
	Create(ctx context.Context, dto domain.CreateTodoDTO) (string, error)
	GetByListId(ctx context.Context, listId, userId string) ([]domain.TodoItem, error)
	GetByUserId(ctx context.Context, userId string) ([]domain.TodoItem, error)
	GetAll(ctx context.Context, userId string) ([]domain.Todo, error)
	GetById(ctx context.Context, itemId string) (domain.TodoItem, error)
	Update(ctx context.Context, dto domain.UpdateTodoDTO) error
	Remove(ctx context.Context, itemId string) error
}

type Services struct {
	Session
	User
	Category
	TodoList
	TodoItem
}

type Deps struct {
	Repos                  *repository.Repositories
	Hasher                 hash.PasswordHasher
	TokenManager           auth.TokenManager
	AccessTokenTTL         time.Duration
	RefreshTokenTTL        time.Duration
	Domain                 string
	VerificationCodeLength int
}

func NewServices(deps Deps) *Services {
	return &Services{
		Session:  NewSessionService(deps.Repos.Users, deps.Repos.Session, deps.TokenManager, deps.Hasher, deps.AccessTokenTTL, deps.RefreshTokenTTL, deps.Domain),
		User:     NewUserService(deps.Repos.Users, deps.TokenManager, deps.Hasher),
		Category: NewCategoryService(deps.Repos.Category, deps.Repos.TodoList, deps.Repos.TodoItem),
		TodoList: NewTodoListService(deps.Repos.TodoList, deps.Repos.TodoItem),
		TodoItem: NewTodoItemService(deps.Repos.TodoItem),
	}
}
