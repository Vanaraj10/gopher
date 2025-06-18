package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Post struct {
    ID        primitive.ObjectID   `bson:"_id,omitempty" json:"id,omitempty"`
    AuthorID  primitive.ObjectID   `bson:"author_id" json:"author_id"`
    Title     string               `bson:"title" json:"title" validate:"required"`
    Content   string               `bson:"content" json:"content" validate:"required"`
    Tags      []string             `bson:"tags,omitempty" json:"tags,omitempty"`
    Likes     []primitive.ObjectID `bson:"likes,omitempty" json:"likes,omitempty"`
    CreatedAt time.Time            `bson:"created_at" json:"created_at"`
    UpdatedAt time.Time            `bson:"updated_at" json:"updated_at"`
}
