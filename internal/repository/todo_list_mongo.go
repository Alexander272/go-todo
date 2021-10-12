package repository

import (
	"context"
	"errors"

	"github.com/Alexander272/go-todo/internal/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type TodoListRepo struct {
	db *mongo.Collection
}

func NewTodoListRepo(db *mongo.Database) *TodoListRepo {
	return &TodoListRepo{
		db: db.Collection(todoListCollection),
	}
}

func (r *TodoListRepo) Create(ctx context.Context, list domain.TodoList) error {
	_, err := r.db.InsertOne(ctx, list)
	return err
}

func (r *TodoListRepo) GetAll(ctx context.Context, userId primitive.ObjectID) ([]domain.TodoList, error) {
	cursor, err := r.db.Find(ctx, bson.M{"userId": userId})
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, domain.ErrListNotFound
		}
		return nil, err
	}

	var lists []domain.TodoList
	if err := cursor.All(ctx, &lists); err != nil {
		return nil, err
	}
	return lists, nil
}

func (r *TodoListRepo) GetById(ctx context.Context, listId primitive.ObjectID) (*domain.TodoList, error) {
	var list *domain.TodoList
	if err := r.db.FindOne(ctx, bson.M{"_id": listId}).Decode(&list); err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, domain.ErrListNotFound
		}

		return nil, err
	}

	return list, nil
}

func (r *TodoListRepo) GetByTitle(ctx context.Context, userId primitive.ObjectID, title string) (*domain.TodoList, error) {
	var list *domain.TodoList
	if err := r.db.FindOne(ctx, bson.M{"title": title, "userId": userId}).Decode(&list); err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, domain.ErrListNotFound
		}

		return nil, err
	}

	return list, nil
}

func (r *TodoListRepo) Update(ctx context.Context, input domain.TodoList) error {
	update := bson.M{}
	if input.Title != "" {
		update["title"] = input.Title
	}
	if input.Description != "" {
		update["description"] = input.Description
	}

	_, err := r.db.UpdateOne(ctx, bson.M{"_id": input.Id}, bson.M{"$set": update})
	return err
}

func (r *TodoListRepo) Remove(ctx context.Context, listId primitive.ObjectID) error {
	_, err := r.db.DeleteOne(ctx, bson.M{"_id": listId})
	return err
}
