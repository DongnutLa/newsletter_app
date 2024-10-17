package domain

type CreateNewsletterDTO struct {
	Template   string   `json:"template"`
	File       string   `json:"file"`
	Recipients []string `json:"recipients"`
}
