package domain

type SendNewsletterDTO struct {
	NewsletterId string `json:"newsletterId"`
	ExtraEmail   string `json:"extraEmail"`
}
