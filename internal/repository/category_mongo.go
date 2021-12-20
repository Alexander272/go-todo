package repository

import (
	"context"
	"errors"
	"fmt"

	"github.com/Alexander272/go-todo/internal/domain"
	"github.com/Alexander272/go-todo/pkg/logger"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type CategoryRepo struct {
	db *mongo.Collection
}

func NewCategoryRepo(db *mongo.Database) *CategoryRepo {
	return &CategoryRepo{
		db: db.Collection(categoryCollection),
	}
}

func (r *CategoryRepo) Create(ctx context.Context, cat domain.Category) (id string, err error) {
	res, err := r.db.InsertOne(ctx, cat)
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

func (r *CategoryRepo) GetAll(ctx context.Context, userId string) (categories []domain.Category, err error) {
	filter := bson.M{"userId": userId}
	cursor, err := r.db.Find(ctx, filter)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return categories, domain.ErrCategoryNotFound
		}
		return categories, fmt.Errorf("failed to execute query. error: %w", err)
	}

	if err := cursor.All(ctx, &categories); err != nil {
		return categories, fmt.Errorf("failed to decode document. error: %w", err)
	}
	return categories, nil
}

func (r *CategoryRepo) GetWithLists(ctx context.Context, userId string) (categories []domain.CategoryWithLists, err error) {
	// filter := bson.M{"userId": userId}
	pipeline := []bson.M{
		{"$match": bson.M{"userId": userId}},
		{"$addFields": bson.M{"categoryId": bson.M{"$toString": "$_id"}}},
		{
			"$lookup": bson.M{
				"from":         "todoList",
				"localField":   "categoryId",
				"foreignField": "categoryId",
				"as":           "lists",
			},
		},
	}
	cursor, err := r.db.Aggregate(ctx, pipeline)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return categories, domain.ErrCategoryNotFound
		}
		return categories, fmt.Errorf("failed to execute query. error: %w", err)
	}

	if err := cursor.All(ctx, &categories); err != nil {
		return categories, fmt.Errorf("failed to decode document. error: %w", err)
	}
	return categories, nil
}

func (r *CategoryRepo) GetByTitle(ctx context.Context, userId, title string) (c domain.Category, err error) {
	filter := bson.M{"title": title, "userId": userId}
	res := r.db.FindOne(ctx, filter)
	if res.Err() != nil {
		if errors.Is(res.Err(), mongo.ErrNoDocuments) {
			return c, domain.ErrCategoryNotFound
		}
		return c, fmt.Errorf("failed to execute query. error: %w", res.Err())
	}
	if err := res.Decode(&c); err != nil {
		return c, fmt.Errorf("failed to decode document. error: %w", err)
	}

	return c, nil
}

func (r *CategoryRepo) Update(ctx context.Context, category domain.Category) error {
	oid, err := primitive.ObjectIDFromHex(category.Id)
	if err != nil {
		return fmt.Errorf("failed to convert hex to objectid. error: %w", err)
	}

	filter := bson.M{"_id": oid}
	categoryByte, err := bson.Marshal(category)
	if err != nil {
		return fmt.Errorf("failed to marshal document. error: %w", err)
	}

	var updateObj bson.M
	if err := bson.Unmarshal(categoryByte, &updateObj); err != nil {
		return fmt.Errorf("failed to unmarshal document. error: %w", err)
	}

	delete(updateObj, "_id")
	update := bson.M{"$set": updateObj}

	res, err := r.db.UpdateOne(ctx, filter, update)
	if err != nil {
		return fmt.Errorf("failed to execute query. error: %w", err)
	}
	if res.MatchedCount == 0 {
		return domain.ErrCategoryNotFound
	}

	logger.Tracef("Matched %v documents and updated %v documents.\n", res.MatchedCount, res.ModifiedCount)
	return nil
}

func (r *CategoryRepo) Remove(ctx context.Context, categoryId string) error {
	oid, err := primitive.ObjectIDFromHex(categoryId)
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
