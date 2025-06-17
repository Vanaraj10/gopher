package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// User represents a user account in the system.
// It contains identification, authentication, and profile information.
// Fields:
//   - ID: Unique identifier for the user (MongoDB ObjectID).
//   - Name: The user's display name (required, 3-30 characters).
//   - Email: The user's email address (required, must be valid).
//   - Password: The user's hashed password (required, 6-100 characters).
//   - Verified: Indicates whether the user's email is verified.
//   - CreatedAt: Timestamp of when the user was created.
type User struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Name      string             `bson:"name" json:"name" validate:"required,min=3,max=30"`
	Email     string             `bson:"email" json:"email" validate:"required,email"`
	Password  string             `bson:"password" json:"password" validate:"required,min=6,max=100"`
	Verified  bool               `bson:"verified" json:"verified"`
	CreatedAt time.Time          `bson:"created_at" json:"created_at"`
}
