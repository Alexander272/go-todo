package repository

import (
	"context"

	"github.com/Alexander272/go-todo/internal/domain"
	"github.com/go-redis/redis/v8"
	"go.mongodb.org/mongo-driver/mongo"
)

type Users interface {
	Create(ctx context.Context, user domain.User) (string, error)
	GetAll(ctx context.Context) ([]domain.User, error)
	Verify(ctx context.Context, userId string, code string) error
	SetSession(ctx context.Context, userId string) error
	GetByEmail(ctx context.Context, email string) (domain.User, error)
	GetById(ctx context.Context, userId string) (domain.User, error)
	Update(ctx context.Context, user domain.User) error
	Remove(ctx context.Context, userId string) error
}

type Session interface {
	CreateSession(ctx context.Context, token string, data SessionData) error
	GetDelSession(ctx context.Context, token string) (SessionData, error)
	RemoveSession(ctx context.Context, token string) error
}

type Category interface {
	Create(ctx context.Context, cat domain.Category) (string, error)
	GetAll(ctx context.Context, userId string) ([]domain.Category, error)
	GetWithLists(ctx context.Context, userId string) ([]domain.CategoryWithLists, error)
	GetByTitle(ctx context.Context, userId, title string) (domain.Category, error)
	Update(ctx context.Context, category domain.Category) error
	Remove(ctx context.Context, categoryId string) error
}

type TodoList interface {
	Create(ctx context.Context, list domain.TodoList) (string, error)
	GetAll(ctx context.Context, userId string) ([]domain.TodoList, error)
	GetById(ctx context.Context, listId string) (domain.TodoList, error)
	GetByTitle(ctx context.Context, userId string, categoryId, title string) (domain.TodoList, error)
	RemoveCatogoryId(ctx context.Context, categoryId string) error
	Update(ctx context.Context, list domain.TodoList) error
	Remove(ctx context.Context, listId string) error
}

type TodoItem interface {
	Create(ctx context.Context, item domain.TodoItem) (string, error)
	GetByListId(ctx context.Context, listId, userId string) ([]domain.TodoItem, error)
	GetByUserId(ctx context.Context, userId string) ([]domain.TodoItem, error)
	GetAll(ctx context.Context, userId string) ([]domain.Todo, error)
	GetById(ctx context.Context, itemId string) (domain.TodoItem, error)
	GetByTitle(ctx context.Context, listId string, title string) (domain.TodoItem, error)
	Update(ctx context.Context, item domain.TodoItem) error
	Remove(ctx context.Context, itemId string) error
	RemoveByListId(ctx context.Context, listId string) error
}

type Repositories struct {
	Users
	Session
	Category
	TodoList
	TodoItem
}

func NewRepositories(db *mongo.Database, client redis.Cmdable) *Repositories {
	return &Repositories{
		Session:  NewSessionRepo(client),
		Users:    NewUsersRepo(db),
		Category: NewCategoryRepo(db),
		TodoList: NewTodoListRepo(db),
		TodoItem: NewTodoItemRepo(db),
	}
}
