package model

type LasagnaLoveVolunteerInfo struct {
	Birthday                     string   `json:"birthday,omitempty"`
	GenderIdentity               string   `json:"gender_identity,omitempty"`
	VolunteeringWith             string   `json:"volunteering_with,omitempty"`
	Employer                     string   `json:"employer,omitempty"`
	FacebookName                 string   `json:"facebook_name,omitempty"`
	MaxTravelDistance            int      `json:"max_travel_distance"`
	FamiliesPerDelivery          int      `json:"families_per_delivery"`
	AllowableDietaryRestrictions []string `json:"allowable_dietary_restrictions,omitempty"`
	AccomodatesExtraRequests     bool     `json:"accomodates_extra_requests,omitempty"`
	ShowCompletedRequests        bool     `json:"show_completed_requests,omitempty"`
	AvailableSchedule            []string `json:"available_schedule,omitempty"`
}
