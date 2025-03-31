package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type BookSchema struct {
	ID        primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Title     string             `json:"title" bson:"title"`
	Author    string             `json:"author" bson:"author"`
	Year      int                `json:"year" bson:"year"`
	CreatedAt time.Time          `json:"createdAt" bson:"createdAt"`
	UpdatedAt time.Time          `json:"updatedAt" bson:"updatedAt"`
}
