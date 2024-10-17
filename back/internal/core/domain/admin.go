package domain

import "go.mongodb.org/mongo-driver/bson/primitive"

type Admin struct {
	ID       primitive.ObjectID `json:"id" bson:"_id"`
	Email    string             `json:"email" bson:"email"`
	Name     string             `json:"name" bson:"name"`
	Password string             `json:"password" bson:"password"`
	Token    string             `json:"token" bson:"-"`
}

func NewAdmin(name, email string) *Admin {
	id := primitive.NewObjectID()

	return &Admin{
		ID:    id,
		Name:  name,
		Email: email,
	}
}
