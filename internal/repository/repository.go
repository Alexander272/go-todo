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

type Repositories struct {
	Users
	Auth
}

func NewRepositories(db *mongo.Database, client *redis.Client) *Repositories {
	return &Repositories{
		Auth:  NewAuthRepo(client),
		Users: NewUsersRepo(db),
	}
}
