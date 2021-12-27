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

type TodoItemRepo struct {
	db *mongo.Collection
}

func NewTodoItemRepo(db *mongo.Database) *TodoItemRepo {
	return &TodoItemRepo{
		db: db.Collection(todoItemCollection),
	}
}

func (r *TodoItemRepo) Create(ctx context.Context, item domain.TodoItem) (id string, err error) {
	item.Done = false
	item.CreatedAt = time.Now().Unix()

	res, err := r.db.InsertOne(ctx, item)
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

func (r *TodoItemRepo) GetByListId(ctx context.Context, listId, userId string) (items []domain.TodoItem, err error) {
	filter := bson.M{"listId": listId, "userId": userId}
	cursor, err := r.db.Find(ctx, filter)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return items, domain.ErrItemNotFound
		}
		return items, fmt.Errorf("failed to execute query. error: %w", err)
	}

	if err := cursor.All(ctx, &items); err != nil {
		return items, fmt.Errorf("failed to decode document. error: %w", err)
	}
	return items, nil
}

// а нужен ли этот запрос?
func (r *TodoItemRepo) GetByUserId(ctx context.Context, userId string) (items []domain.TodoItem, err error) {
	filter := bson.M{"userId": userId}
	cursor, err := r.db.Find(ctx, filter)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return items, domain.ErrItemNotFound
		}
		return items, fmt.Errorf("failed to execute query. error: %w", err)
	}

	if err := cursor.All(ctx, &items); err != nil {
		return items, fmt.Errorf("failed to decode document. error: %w", err)
	}
	return items, nil
}

func (r *TodoItemRepo) GetAll(ctx context.Context, userId string) (items []domain.Todo, err error) {
	pipeline := []bson.M{
		{"$match": bson.M{"userId": userId}},
		{"$group": bson.M{
			"_id": "$listId",
			"items": bson.M{"$push": bson.M{
				"_id":   "$_id",
				"title": "$title",
				"done":  "$done",
			}},
		}},
	}

	cursor, err := r.db.Aggregate(ctx, pipeline)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return items, domain.ErrListNotFound
		}
		return items, fmt.Errorf("failed to execute query. error: %w", err)
	}

	if err := cursor.All(ctx, &items); err != nil {
		return items, fmt.Errorf("failed to decode document. error: %w", err)
	}
	return items, nil
}

func (r *TodoItemRepo) GetById(ctx context.Context, itemId string) (item domain.TodoItem, err error) {
	oid, err := primitive.ObjectIDFromHex(itemId)
	if err != nil {
		return item, fmt.Errorf("failed to convert hex to objectid. error: %w", err)
	}

	filter := bson.M{"_id": oid}
	res := r.db.FindOne(ctx, filter)
	if res.Err() != nil {
		if errors.Is(res.Err(), mongo.ErrNoDocuments) {
			return item, domain.ErrItemNotFound
		}
		return item, fmt.Errorf("failed to execute query. error: %w", res.Err())
	}
	if err := res.Decode(&item); err != nil {
		return item, fmt.Errorf("failed to decode document. error: %w", err)
	}

	return item, nil
}

func (r *TodoItemRepo) GetByTitle(ctx context.Context, listId string, title string) (item domain.TodoItem, err error) {
	filter := bson.M{"listId": listId, "title": title}
	res := r.db.FindOne(ctx, filter)
	if res.Err() != nil {
		if errors.Is(res.Err(), mongo.ErrNoDocuments) {
			return item, domain.ErrItemNotFound
		}
		return item, fmt.Errorf("failed to execute query. error: %w", res.Err())
	}
	if err := res.Decode(&item); err != nil {
		return item, fmt.Errorf("failed to decode document. error: %w", err)
	}

	return item, nil
}

func (r *TodoItemRepo) Update(ctx context.Context, item domain.TodoItem) error {
	oid, err := primitive.ObjectIDFromHex(item.Id)
	if err != nil {
		return fmt.Errorf("failed to convert hex to objectid. error: %w", err)
	}

	filter := bson.M{"_id": oid}
	itemByte, err := bson.Marshal(item)
	if err != nil {
		return fmt.Errorf("failed to marshal document. error: %w", err)
	}

	var updateObj bson.M
	if err := bson.Unmarshal(itemByte, &updateObj); err != nil {
		return fmt.Errorf("failed to unmarshal document. error: %w", err)
	}

	delete(updateObj, "_id")
	update := bson.M{"$set": updateObj}

	res, err := r.db.UpdateOne(ctx, filter, update)
	if err != nil {
		return fmt.Errorf("failed to execute query. error: %w", err)
	}
	if res.MatchedCount == 0 {
		return domain.ErrItemNotFound
	}

	logger.Tracef("Matched %v documents and updated %v documents.\n", res.MatchedCount, res.ModifiedCount)
	return nil
}

func (r *TodoItemRepo) Remove(ctx context.Context, itemId string) error {
	oid, err := primitive.ObjectIDFromHex(itemId)
	if err != nil {
		return fmt.Errorf("failed to convert hex to objectid. error: %w", err)
	}

	filter := bson.M{"_id": oid}
	res, err := r.db.DeleteOne(ctx, filter)
	if err != nil {
		return fmt.Errorf("failed to execute query. error: %w", err)
	}
	if res.DeletedCount == 0 {
		return domain.ErrItemNotFound
	}

	logger.Tracef("Delete %v documents.\n", res.DeletedCount)
	return nil
}

func (r *TodoItemRepo) RemoveByListId(ctx context.Context, listId string) error {
	filter := bson.M{"listId": listId}
	res, err := r.db.DeleteMany(ctx, filter)
	if err != nil {
		return fmt.Errorf("failed to execute query. error: %w", err)
	}
	if res.DeletedCount == 0 {
		return domain.ErrItemNotFound
	}

	logger.Tracef("Delete %v documents.\n", res.DeletedCount)
	return nil
}
