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

type UsersRepo struct {
	db *mongo.Collection
}

func NewUsersRepo(db *mongo.Database) *UsersRepo {
	return &UsersRepo{
		db: db.Collection(usersCollection),
	}
}

func (r *UsersRepo) Create(ctx context.Context, user domain.User) (id string, err error) {
	user.RegisteredAt = time.Now().Unix()
	res, err := r.db.InsertOne(ctx, user)
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

func (r *UsersRepo) GetAll(ctx context.Context) (users []domain.User, err error) {
	filter := bson.M{}
	cur, err := r.db.Find(ctx, filter)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return users, domain.ErrUserNotFound
		}
		return users, fmt.Errorf("failed to execute query. error: %w", err)
	}
	if err := cur.All(ctx, &users); err != nil {
		return users, fmt.Errorf("failed to decode document. error: %w", err)
	}

	return users, nil
}

func (r *UsersRepo) Verify(ctx context.Context, userId, code string) error {
	res, err := r.db.UpdateOne(ctx,
		bson.M{"verification.code": code, "_id": userId},
		bson.M{"$set": bson.M{"verification.verified": true, "verification.code": ""}})
	if err != nil {
		return err
	}

	if res.ModifiedCount == 0 {
		return domain.ErrVerificationCodeInvalid
	}

	return nil
}

func (r *UsersRepo) SetSession(ctx context.Context, userId string) error {
	oid, err := primitive.ObjectIDFromHex(userId)
	if err != nil {
		return fmt.Errorf("failed to convert hex to objectid. error: %w", err)
	}

	filter := bson.M{"_id": oid}
	updateObj := bson.M{"$set": bson.M{"lastVisitAt": time.Now().Unix()}}

	res, err := r.db.UpdateOne(ctx, filter, updateObj)
	if err != nil {
		return fmt.Errorf("failed to execute query. error: %w", err)
	}
	if res.MatchedCount == 0 {
		return domain.ErrUserNotFound
	}

	logger.Tracef("Matched %v documents and updated %v documents.\n", res.MatchedCount, res.ModifiedCount)
	return nil
}

func (r *UsersRepo) GetByEmail(ctx context.Context, email string) (user domain.User, err error) {
	filter := bson.M{"email": email}
	res := r.db.FindOne(ctx, filter)
	if res.Err() != nil {
		if errors.Is(res.Err(), mongo.ErrNoDocuments) {
			return user, domain.ErrUserNotFound
		}
		return user, fmt.Errorf("failed to execute query. error: %w", res.Err())
	}
	if err := res.Decode(&user); err != nil {
		return user, fmt.Errorf("failed to decode document. error: %w", err)
	}

	return user, nil
}

func (r *UsersRepo) GetById(ctx context.Context, userId string) (user domain.User, err error) {
	oid, err := primitive.ObjectIDFromHex(userId)
	if err != nil {
		return user, fmt.Errorf("failed to convert hex to objectid. error: %w", err)
	}

	filter := bson.M{"_id": oid}
	res := r.db.FindOne(ctx, filter)
	if res.Err() != nil {
		if errors.Is(res.Err(), mongo.ErrNoDocuments) {
			return user, domain.ErrUserNotFound
		}
		return user, fmt.Errorf("failed to execute query. error: %w", res.Err())
	}
	if err := res.Decode(&user); err != nil {
		return user, fmt.Errorf("failed to decode document. error: %w", err)
	}

	return user, nil
}

func (r *UsersRepo) Update(ctx context.Context, user domain.User) error {
	oid, err := primitive.ObjectIDFromHex(user.Id)
	if err != nil {
		return fmt.Errorf("failed to convert hex to objectid. error: %w", err)
	}

	filter := bson.M{"_id": oid}
	userByte, err := bson.Marshal(user)
	if err != nil {
		return fmt.Errorf("failed to marshal document. error: %w", err)
	}

	var updateObj bson.M
	if err := bson.Unmarshal(userByte, &updateObj); err != nil {
		return fmt.Errorf("failed to unmarshal document. error: %w", err)
	}

	delete(updateObj, "_id")
	update := bson.M{"$set": updateObj}

	res, err := r.db.UpdateOne(ctx, filter, update)
	if err != nil {
		return fmt.Errorf("failed to execute query. error: %w", err)
	}
	if res.MatchedCount == 0 {
		return domain.ErrUserNotFound
	}

	logger.Tracef("Matched %v documents and updated %v documents.\n", res.MatchedCount, res.ModifiedCount)
	return nil
}

func (r *UsersRepo) Remove(ctx context.Context, userId string) error {
	oid, err := primitive.ObjectIDFromHex(userId)
	if err != nil {
		return fmt.Errorf("failed to convert hex to objectid. error: %w", err)
	}

	filter := bson.M{"_id": oid}
	res, err := r.db.DeleteOne(ctx, filter)
	if err != nil {
		return fmt.Errorf("failed to execute query. error: %w", err)
	}
	if res.DeletedCount == 0 {
		return domain.ErrUserNotFound
	}

	logger.Tracef("Delete %v documents.\n", res.DeletedCount)
	return nil
}
