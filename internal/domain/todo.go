package domain

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

//todo Можно добавить возможность делиться задачами и листами

type TodoList struct {
	Id          primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	UserId      primitive.ObjectID `json:"userId" bson:"userId,omitempty"`
	Title       string             `json:"title" bson:"title"`
	Description string             `json:"description" bson:"description"`
	CreatedAt   time.Time          `json:"createdAt" bson:"createdAt"`
}

type TodoItem struct {
	Id          primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	UserId      primitive.ObjectID `json:"userId" bson:"userId,omitempty"`
	ListId      primitive.ObjectID `json:"listId" bson:"listId,omitempty"`
	Title       string             `json:"title" bson:"title"`
	Description string             `json:"description" bson:"description"`
	CreatedAt   time.Time          `json:"createdAt" bson:"createdAt"`
	CompletedAt time.Time          `json:"completedAt" bson:"completedAt"`
	DeadlineAt  time.Time          `json:"deadlineAt" bson:"deadlineAt"`
	Done        bool               `json:"done" bson:"done"`
	Priority    int                `json:"priority" bson:"priority"`
	Tags        []string           `json:"tags" bson:"tags"`
	// todo можно потом добавить Files File
}
