package domain

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID        primitive.ObjectID `bson:"_id" json:"id"`
	Email     string             `bson:"email" json:"email"`
	Topics    []string           `bson:"topics" json:"topics"`
	Timestamp *time.Time         `bson:"timestamp" json:"timestamp"`
}

func NewUser(email string, topics []string) *User {
	id := primitive.NewObjectID()
	now := time.Now()

	return &User{
		ID:        id,
		Email:     email,
		Topics:    topics,
		Timestamp: &now,
	}
}
