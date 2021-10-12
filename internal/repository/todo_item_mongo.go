package repository

import (
	"context"
	"errors"
	"time"

	"github.com/Alexander272/go-todo/internal/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type TodoItemRepo struct {
	db *mongo.Collection
}

func NewTodoItemRepo(db *mongo.Database) *TodoItemRepo {
	return &TodoItemRepo{
		db: db.Collection(todoItemCollection),
	}
}

func (r *TodoItemRepo) Create(ctx context.Context, item domain.TodoItem) error {
	_, err := r.db.InsertOne(ctx, item)
	return err
}

func (r *TodoItemRepo) GetByListId(ctx context.Context, userId primitive.ObjectID, listId primitive.ObjectID) ([]domain.TodoItem, error) {
	cursor, err := r.db.Find(ctx, bson.M{"userId": userId, "listId": listId})
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, domain.ErrItemNotFound
		}
		return nil, err
	}

	var items []domain.TodoItem
	if err := cursor.All(ctx, &items); err != nil {
		return nil, err
	}
	return items, nil
}

func (r *TodoItemRepo) GetByUserId(ctx context.Context, userId primitive.ObjectID) ([]domain.TodoItem, error) {
	cursor, err := r.db.Find(ctx, bson.M{"userId": userId})
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, domain.ErrItemNotFound
		}
		return nil, err
	}

	var items []domain.TodoItem
	if err := cursor.All(ctx, &items); err != nil {
		return nil, err
	}
	return items, nil
}

func (r *TodoItemRepo) GetById(ctx context.Context, itemId primitive.ObjectID) (*domain.TodoItem, error) {
	var item *domain.TodoItem
	if err := r.db.FindOne(ctx, bson.M{"_id": itemId}).Decode(&item); err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, domain.ErrItemNotFound
		}
		return nil, err
	}

	return item, nil
}

func (r *TodoItemRepo) GetByTitle(ctx context.Context, userId primitive.ObjectID, title string) (*domain.TodoItem, error) {
	var item *domain.TodoItem
	if err := r.db.FindOne(ctx, bson.M{"userId": userId, "title": title}).Decode(&item); err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, domain.ErrItemNotFound
		}
		return nil, err
	}

	return item, nil
}

func (r *TodoItemRepo) Update(ctx context.Context, input domain.TodoItem) error {
	update := bson.M{}
	if input.Title != "" {
		update["title"] = input.Title
	}
	if input.Description != "" {
		update["description"] = input.Description
	}
	if input.ListId != primitive.NilObjectID {
		update["listId"] = input.ListId
	}
	if input.Priority != 0 {
		update["priority"] = input.Priority
	}
	if !input.DeadlineAt.IsZero() {
		update["deadlineAt"] = input.DeadlineAt
	}
	if input.Done {
		update["completedAt"] = time.Now()
	}
	if len(input.Tags) != 0 {
		update["tags"] = input.Tags
	}
	update["done"] = input.Done

	_, err := r.db.UpdateOne(ctx, bson.M{"_id": input.Id}, bson.M{"$set": update})
	return err
}

func (r *TodoItemRepo) Remove(ctx context.Context, itemId primitive.ObjectID) error {
	_, err := r.db.DeleteOne(ctx, bson.M{"_id": itemId})
	return err
}
