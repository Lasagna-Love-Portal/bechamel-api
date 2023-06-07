package model

type LasagnaLoveUser struct {
	ID                      int                       `json:"id"`
	Roles                   []string                  `json:"roles"`
	Username                string                    `json:"username"`
	Password                string                    `json:"-"`
	GivenName               string                    `json:"given_name"`
	MiddleOrMaidenName      string                    `json:"middle_or_maiden_name,omitempty"`
	FamilyName              string                    `json:"family_name"`
	Email                   string                    `json:"email"`
	EmailValidated          bool                      `json:"email_validated"`
	CreationTime            string                    `json:"creation_time"`
	LastUpdateTime          string                    `json:"last_update_time"`
	StreetAddress           []string                  `json:"street_address"`
	City                    string                    `json:"city"`
	StateOrProvince         string                    `json:"state_or_province"`
	PostalCode              string                    `json:"postal_code"`
	HomePhone               string                    `json:"home_phone,omitempty"`
	MobilePhone             string                    `json:"mobile_phone"`
	MobileContactPermission bool                      `json:"mobile_contact_permission"`
	NewsUpdatesPermission   bool                      `json:"news_updates_permission"`
	Active                  bool                      `json:"active"`
	Paused                  bool                      `json:"paused"`
	PausedEndDate           string                    `json:"paused_end_date,omitempty"`
	Attestations            LasagnaLoveAttestations   `json:"attestations"`
	RecipientInfo           *LasagnaLoveRecipientInfo `json:"recipient_info,omitempty"`
	VolunteerInfo           *LasagnaLoveVolunteerInfo `json:"volunteer_info,omitempty"`
}
