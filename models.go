package main

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"gorm.io/gorm"
)

// mongoDB
type UserDetails struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	Name      string             `bson:"title,omitempty"`
	Age       int16              `bson:"author,omitempty"`
	Email     string             `bson:"email,omitempty"`
	Active    bool               `bson:"active,omitempty"`
	CreatedAt time.Time          `bson:"created_at"`
	UpdatedAt time.Time          `bson:"updated_at"`
}

// postgres
// ID, createdAt, updatedAt are automatically added by gorm
type User struct {
	gorm.Model
	Name   string
	Email  string `gorm:"uniqueIndex"`
	Age    uint8
	Active bool `gorm:"default:true"`
}
