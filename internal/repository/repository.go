package repository

import (
	"context"

	"github.com/Alexander272/go-todo/internal/domain"
	"github.com/go-redis/redis/v8"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Users interface {
	Create(ctx context.Context, user domain.User) error
	GetByEmail(ctx context.Context, email string) (*domain.User, error)
	Verify(ctx context.Context, userId primitive.ObjectID, code string) error
	SetSession(ctx context.Context, userId primitive.ObjectID) error
	GetById(ctx context.Context, userId primitive.ObjectID) (domain.User, error)
	UpdateById(ctx context.Context, userId primitive.ObjectID, user domain.UserUpdate) error
	RemoveById(ctx context.Context, userId primitive.ObjectID) error
	GetAllUsers(ctx context.Context) ([]domain.User, error)
}

type Auth interface {
	CreateSession(token string, data RedisData) error
	GetDelSession(token string) (*RedisData, error)
	RemoveSession(token string) error
}

type TodoList interface {
	Create(ctx context.Context, list domain.TodoList) error
	GetAll(ctx context.Context, userId primitive.ObjectID) ([]domain.TodoList, error)
	GetAllWithTodo(ctx context.Context, userId primitive.ObjectID) ([]domain.TodoListWithItems, error)
	GetById(ctx context.Context, listId primitive.ObjectID) (*domain.TodoList, error)
	GetByTitle(ctx context.Context, userId primitive.ObjectID, title string) (*domain.TodoList, error)
	Update(ctx context.Context, input domain.TodoList) error
	Remove(ctx context.Context, listId primitive.ObjectID) error
}

type TodoItem interface {
	Create(ctx context.Context, item domain.TodoItem) error
	GetByListId(ctx context.Context, userId primitive.ObjectID, listId primitive.ObjectID) ([]domain.TodoItem, error)
	GetByUserId(ctx context.Context, userId primitive.ObjectID) ([]domain.TodoItem, error)
	GetById(ctx context.Context, itemId primitive.ObjectID) (*domain.TodoItem, error)
	GetByTitle(ctx context.Context, userId primitive.ObjectID, title string) (*domain.TodoItem, error)
	Update(ctx context.Context, input domain.TodoItem) error
	Remove(ctx context.Context, itemId primitive.ObjectID) error
	RemoveByListId(ctx context.Context, listId primitive.ObjectID) error
}

type Repositories struct {
	Users
	Auth
	TodoList
	TodoItem
}

func NewRepositories(db *mongo.Database, client *redis.Client) *Repositories {
	return &Repositories{
		Auth:     NewAuthRepo(client),
		Users:    NewUsersRepo(db),
		TodoList: NewTodoListRepo(db),
		TodoItem: NewTodoItemRepo(db),
	}
}
