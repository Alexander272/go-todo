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
	CreatedAt   int64              `json:"createdAt" bson:"createdAt"`
}

type TodoListWithItems struct {
	Id          primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	UserId      primitive.ObjectID `json:"userId" bson:"userId,omitempty"`
	Title       string             `json:"title" bson:"title"`
	Description string             `json:"description" bson:"description"`
	CreatedAt   int64              `json:"createdAt" bson:"createdAt"`
	Todos       []Todo             `json:"todos" bson:"todos"`
}

type Todo struct {
	Id         primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Title      string             `json:"title" bson:"title"`
	DeadlineAt time.Time          `json:"deadlineAt" bson:"deadlineAt,omitempty"`
	Done       bool               `json:"done" bson:"done"`
	Priority   int                `json:"priority" bson:"priority"`
}

type TodoItem struct {
	Id          primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	UserId      primitive.ObjectID `json:"userId" bson:"userId,omitempty"`
	ListId      primitive.ObjectID `json:"listId" bson:"listId,omitempty"`
	Title       string             `json:"title" bson:"title"`
	Description string             `json:"description" bson:"description"`
	CreatedAt   int64              `json:"createdAt" bson:"createdAt"`
	CompletedAt int64              `json:"completedAt" bson:"completedAt,omitempty"`
	DeadlineAt  int64              `json:"deadlineAt" bson:"deadlineAt,omitempty"`
	Done        bool               `json:"done" bson:"done"`
	Priority    int                `json:"priority" bson:"priority"`
	Tags        []string           `json:"tags" bson:"tags"`
	// todo можно потом добавить Files File
}
