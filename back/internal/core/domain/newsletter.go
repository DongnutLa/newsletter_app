package domain

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Newsletter struct {
	ID         primitive.ObjectID `json:"id" bson:"_id"`
	Template   string             `json:"template" bson:"template"`
	Subject    string             `json:"subject" bson:"subject"`
	File       string             `json:"file" bson:"file"`
	Recipients []string           `json:"recipients" bson:"recipients"`
	Topic      string             `json:"topic" bson:"topic"`
	Active     bool               `json:"active" bson:"active"`
	Schedule   string             `json:"schedule" bson:"schedule"`
	SentCount  uint64             `json:"SentCount" bson:"SentCount"`
	Timestamp  *time.Time         `json:"timestamp" bson:"timestamp"`
}

func NewNewsletter(dto *CreateNewsletterDTO) *Newsletter {
	id := primitive.NewObjectID()
	now := time.Now()

	return &Newsletter{
		ID:         id,
		Template:   dto.Template,
		File:       dto.File,
		Recipients: dto.Recipients,
		Subject:    dto.Subject,
		Topic:      dto.Topic,
		Timestamp:  &now,
	}
}
