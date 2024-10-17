package domain

import "go.mongodb.org/mongo-driver/bson/primitive"

type Newsletter struct {
	ID         string   `json:"id" bson:"_id"`
	Template   string   `json:"template" bson:"template"`
	Subject    string   `json:"subject" bson:"subject"`
	File       string   `json:"file" bson:"file"`
	Recipients []string `json:"recipients" bson:"recipients"`
	Active     bool     `json:"active" bson:"active"`
	Schedule   string   `json:"schedule" bson:"schedule"`
}

func NewNewsletter(template, file string, recipients []string) *Newsletter {
	id := primitive.NewObjectID()

	return &Newsletter{
		ID:         id.Hex(),
		Template:   template,
		File:       file,
		Recipients: recipients,
	}
}
