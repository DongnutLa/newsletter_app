package domain

import "go.mongodb.org/mongo-driver/bson/primitive"

type Admin struct {
	ID       string `json:"id" bson:"_id"`
	Email    string `json:"email" bson:"email"`
	Name     string `json:"name" bson:"name"`
	Password string `json:"password" bson:"password"`
}

func NewAdmin(name, email string) *Admin {
	id := primitive.NewObjectID()

	return &Admin{
		ID:    id.Hex(),
		Name:  name,
		Email: email,
	}
}
