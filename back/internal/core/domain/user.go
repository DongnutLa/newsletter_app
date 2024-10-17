package domain

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	ID    string `bson:"_id" json:"id"`
	Email string `bson:"email" json:"email"`
}

func NewUser(email string) *User {
	id := primitive.NewObjectID()

	return &User{
		ID:    id.Hex(),
		Email: email,
	}
}
