package main

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserDetails struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	Name      string             `bson:"title,omitempty"`
	Age       int16              `bson:"author,omitempty"`
	Email     string             `bson:"email,omitempty"`
	Active    bool               `bson:"active,omitempty"`
	CreatedAt time.Time          `bson:"created_at"`
	UpdatedAt time.Time          `bson:"updated_at"`
}
