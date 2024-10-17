package domain

import "go.mongodb.org/mongo-driver/bson/primitive"

type Newsletter struct {
	ID         primitive.ObjectID `json:"id" bson:"_id"`
	Template   string             `json:"template" bson:"template"`
	Subject    string             `json:"subject" bson:"subject"`
	File       string             `json:"file" bson:"file"`
	Recipients []string           `json:"recipients" bson:"recipients"`
	Topic      string             `json:"topic" bson:"topic"`
	Active     bool               `json:"active" bson:"active"`
	Schedule   string             `json:"schedule" bson:"schedule"`
}

func NewNewsletter(dto *CreateNewsletterDTO) *Newsletter {
	id := primitive.NewObjectID()

	return &Newsletter{
		ID:         id,
		Template:   dto.Template,
		File:       dto.File,
		Recipients: dto.Recipients,
		Subject:    dto.Subject,
		Topic:      dto.Topic,
	}
}
