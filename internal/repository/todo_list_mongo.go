package repository

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/Alexander272/go-todo/internal/domain"
	"github.com/Alexander272/go-todo/pkg/logger"
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

func (r *TodoListRepo) Create(ctx context.Context, list domain.TodoList) (id string, err error) {
	list.CreatedAt = time.Now().Unix()

	res, err := r.db.InsertOne(ctx, list)
	if err != nil {
		return id, fmt.Errorf("failed to execute query. error: %w", err)
	}

	oid, ok := res.InsertedID.(primitive.ObjectID)
	if !ok {
		return id, fmt.Errorf("failed to convert objectid")
	}
	logger.Tracef("Created document with oid %s.\n", oid)
	return oid.Hex(), nil
}

func (r *TodoListRepo) GetAll(ctx context.Context, userId string) (lists []domain.TodoList, err error) {
	filter := bson.M{"userId": userId}
	cursor, err := r.db.Find(ctx, filter)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return lists, domain.ErrListNotFound
		}
		return lists, fmt.Errorf("failed to execute query. error: %w", err)
	}

	if err := cursor.All(ctx, &lists); err != nil {
		return lists, fmt.Errorf("failed to decode document. error: %w", err)
	}
	return lists, nil
}

func (r *TodoListRepo) GetById(ctx context.Context, listId string) (list domain.TodoList, err error) {
	oid, err := primitive.ObjectIDFromHex(listId)
	if err != nil {
		return list, fmt.Errorf("failed to convert hex to objectid. error: %w", err)
	}

	filter := bson.M{"_id": oid}
	res := r.db.FindOne(ctx, filter)
	if res.Err() != nil {
		if errors.Is(res.Err(), mongo.ErrNoDocuments) {
			return list, domain.ErrListNotFound
		}
		return list, fmt.Errorf("failed to execute query. error: %w", res.Err())
	}
	if err := res.Decode(&list); err != nil {
		return list, fmt.Errorf("failed to decode document. error: %w", err)
	}

	return list, nil
}

func (r *TodoListRepo) GetByTitle(ctx context.Context, userId, categoryId, title string) (list domain.TodoList, err error) {
	filter := bson.M{"title": title, "userId": userId, "categoryId": categoryId}
	res := r.db.FindOne(ctx, filter)
	if res.Err() != nil {
		if errors.Is(res.Err(), mongo.ErrNoDocuments) {
			return list, domain.ErrListNotFound
		}
		return list, fmt.Errorf("failed to execute query. error: %w", res.Err())
	}
	if err := res.Decode(&list); err != nil {
		return list, fmt.Errorf("failed to decode document. error: %w", err)
	}

	return list, nil
}

func (r *TodoListRepo) RemoveCatogoryId(ctx context.Context, categoryId string) error {
	filter := bson.M{"categoryId": categoryId}
	upfateObj := bson.M{"categoryId": ""}
	res, err := r.db.UpdateMany(ctx, filter, upfateObj)
	if err != nil {
		return fmt.Errorf("failed to execute query. error: %w", err)
	}
	if res.MatchedCount == 0 {
		return domain.ErrListNotFound
	}

	logger.Tracef("Matched %v documents and updated %v documents.\n", res.MatchedCount, res.ModifiedCount)
	return nil
}

func (r *TodoListRepo) Update(ctx context.Context, todo domain.TodoList) error {
	oid, err := primitive.ObjectIDFromHex(todo.Id)
	if err != nil {
		return fmt.Errorf("failed to convert hex to objectid. error: %w", err)
	}

	filter := bson.M{"_id": oid}
	todoByte, err := bson.Marshal(todo)
	if err != nil {
		return fmt.Errorf("failed to marshal document. error: %w", err)
	}

	var updateObj bson.M
	if err := bson.Unmarshal(todoByte, &updateObj); err != nil {
		return fmt.Errorf("failed to unmarshal document. error: %w", err)
	}

	delete(updateObj, "_id")
	update := bson.M{"$set": updateObj}

	res, err := r.db.UpdateOne(ctx, filter, update)
	if err != nil {
		return fmt.Errorf("failed to execute query. error: %w", err)
	}
	if res.MatchedCount == 0 {
		return domain.ErrListNotFound
	}

	logger.Tracef("Matched %v documents and updated %v documents.\n", res.MatchedCount, res.ModifiedCount)
	return nil
}

func (r *TodoListRepo) Remove(ctx context.Context, listId string) error {
	oid, err := primitive.ObjectIDFromHex(listId)
	if err != nil {
		return fmt.Errorf("failed to convert hex to objectid. error: %w", err)
	}

	filter := bson.M{"_id": oid}
	res, err := r.db.DeleteOne(ctx, filter)
	if err != nil {
		return fmt.Errorf("failed to execute query. error: %w", err)
	}
	if res.DeletedCount == 0 {
		return domain.ErrListNotFound
	}

	logger.Tracef("Delete %v documents.\n", res.DeletedCount)
	return nil
}
