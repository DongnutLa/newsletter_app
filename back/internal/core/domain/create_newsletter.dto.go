package domain

type CreateNewsletterDTO struct {
	Template   string   `json:"template"`
	File       string   `json:"file"`
	Subject    string   `json:"subject"`
	Recipients []string `json:"recipients"`
}
