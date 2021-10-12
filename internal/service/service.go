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
	CreateList(ctx context.Context, input CreateTodoList) error
	GetAllLists(ctx context.Context, userId primitive.ObjectID) ([]domain.TodoList, error)
	GetListById(ctx context.Context, listId primitive.ObjectID) (*domain.TodoList, error)
	UpdateList(ctx context.Context, listId primitive.ObjectID, input UpdateTodolist) error
	RemoveList(ctx context.Context, listId primitive.ObjectID) error
}

type TodoItem interface {
	CreateItem(ctx context.Context, input CreateTodoItem) error
	GetItemsByListId(ctx context.Context, userId, listId primitive.ObjectID) ([]domain.TodoItem, error)
	GetItemsByUserId(ctx context.Context, userId primitive.ObjectID) ([]domain.TodoItem, error)
	GetItemsById(ctx context.Context, itemId primitive.ObjectID) (*domain.TodoItem, error)
	UpdateItem(ctx context.Context, input UpdateTodoItem) error
	RemoveItem(ctx context.Context, itemId primitive.ObjectID) error
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
