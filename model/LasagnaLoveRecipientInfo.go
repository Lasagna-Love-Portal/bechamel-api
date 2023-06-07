package model

type LasagnaLoveRecipientInfo struct {
	AdultCount          int      `json:"adult_count"`
	ChildCount          int      `json:"child_count"`
	DietaryRestrictions []string `json:"dietary_restrictions"`
}
