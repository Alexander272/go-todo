package domain

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

//todo Можно добавить возможность делиться задачами и листами

type Category struct {
	Id     primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	UserId primitive.ObjectID `json:"userId" bson:"userId,omitempty"`
	Title  string             `json:"title" bson:"title"`
}
