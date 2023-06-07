package model

type LasagnaLoveVolunteerInfo struct {
	Birthday                     string   `json:"birthday"`
	GenderIdentity               string   `json:"gender_identity"`
	VolunteeringWith             string   `json:"volunteering_with"`
	Employer                     string   `json:"employer"`
	FacebookName                 string   `json:"facebook_name"`
	MaxTravelDistance            int      `json:"max_travel_distance"`
	AllowableDietaryRestrictions []string `json:"allowable_dietary_restrictions"`
	AccomodatesExtraRequests     bool     `json:"accomodates_extra_requests"`
	ShowCompletedRequests        bool     `json:"show_completed_requests"`
	AvailableSchedule            []string `json:"available_schedule"`
}
