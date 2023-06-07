package model

type LasagnaLoveAttestations struct {
	UserIsEighteen                    bool   `json:"user_is_eighteen"`
	UserAcceptedEmailCommunications   bool   `json:"user_accepted_email_communications"`
	RequesterAcceptedLiabilityRelease string `json:"requester_accepted_liability_release"`
	VolunteerAcceptedIndemnityWaiver  string `json:"volunteer_accepted_indemnity_waiver"`
	VolunteerAcceptedVolunteerTerms   string `json:"volunteer_accepted_volunteer_terms"`
	VolunteerCompletedSafetyTraining  string `json:"volunteer_completed_safety_training"`
}
