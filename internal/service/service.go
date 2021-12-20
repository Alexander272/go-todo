package service

import (
	"context"
	"net/http"
	"time"

	"github.com/Alexander272/go-todo/internal/domain"
	"github.com/Alexander272/go-todo/internal/repository"
	"github.com/Alexander272/go-todo/pkg/auth"
	"github.com/Alexander272/go-todo/pkg/hash"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const CookieName = "session"

type SignUpInput struct {
	Name     string
	Email    string
	Password string
}
type SignInInput struct {
	Email    string
	Password string
}
type Token struct {
	AccessToken  string
	RefreshToken string
}

type Auth interface {
	SignIn(ctx context.Context, input SignInInput, ua, ip string) (*http.Cookie, *Token, error)
	SingOut(token string) (*http.Cookie, error)
	Refresh(token, ua, ip string) (*Token, *http.Cookie, error)
	TokenParse(token string) (userId string, role string, err error)
}

type User interface {
	SignUp(ctx context.Context, input SignUpInput) error
	GetById(ctx context.Context, userId primitive.ObjectID) (domain.User, error)
	UpdateById(ctx context.Context, userId primitive.ObjectID, user domain.UserUpdate) error
	RemoveById(ctx context.Context, userId primitive.ObjectID) error
	GetAllUsers(ctx context.Context) ([]domain.User, error)
}

type TodoList interface {
	CreateList(ctx context.Context, dto domain.CreateListDTO) (string, error)
	GetAllLists(ctx context.Context, userId string) ([]domain.TodoList, error)
	GetListById(ctx context.Context, listId string) (domain.TodoList, error)
	UpdateList(ctx context.Context, listId string, dto domain.UpdateListDTO) error
	RemoveList(ctx context.Context, listId string) error
}

type TodoItem interface {
	Create(ctx context.Context, dto domain.CreateTodoDTO) (string, error)
	GetByListId(ctx context.Context, listId string) ([]domain.TodoItem, error)
	GetByUserId(ctx context.Context, userId string) ([]domain.TodoItem, error)
	GetById(ctx context.Context, itemId string) (domain.TodoItem, error)
	Update(ctx context.Context, dto domain.UpdateTodoDTO) error
	Remove(ctx context.Context, itemId string) error
}

type Services struct {
	Auth
	User
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
		Auth:     NewAuthService(deps.Repos.Users, deps.Repos.Auth, deps.TokenManager, deps.Hasher, deps.AccessTokenTTL, deps.RefreshTokenTTL, deps.Domain),
		User:     NewUserService(deps.Repos.Users, deps.TokenManager, deps.Hasher),
		TodoList: NewTodoListService(deps.Repos.TodoList, deps.Repos.TodoItem),
		TodoItem: NewTodoItemService(deps.Repos.TodoItem),
	}
}
