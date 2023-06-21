package beri_review

type BeriReviewRequest struct {
	PenggunaID string `json:"user_id"`
	PartnerID  string `json:"partner_id"`
	Rating     int    `json:"rating"`
	Comment    string `json:"comment,omitempty"`
}
