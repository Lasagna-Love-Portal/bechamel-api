package model

type LasagnaLoveRecipientInfo struct {
	AdultCount           int      `json:"adult_count"`
	ChildCount           int      `json:"child_count,omitempty"`
	LearnedAboutFrom     string   `json:"learned_about_from"`
	LastDeliveryReceived string   `json:"last_delivery_received"`
	DietaryRestrictions  []string `json:"dietary_restrictions"`
}
