package model

// TODO: Fix up required vs. optional
type LasagnaLoveUser struct {
	ID                      int                      `json:"id"`
	Roles                   []string                 `json:"roles"`
	Username                string                   `json:"username"`
	Password                string                   `json:"-"`
	GivenName               string                   `json:"given_name"`
	MiddleOrMaidenName      string                   `json:"middle_or_maiden_name"`
	FamilyName              string                   `json:"family_name"`
	Email                   string                   `json:"email"`
	EmailValidated          bool                     `json:"email_validated"`
	CreationTime            string                   `json:"creation_time"`
	LastUpdateTime          string                   `json:"last_update_time"`
	StreetAddress           []string                 `json:"street_address"`
	City                    string                   `json:"city"`
	StateOrProvince         string                   `json:"state_or_province"`
	PostalCode              string                   `json:"postal_code"`
	HomePhone               string                   `json:"home_phone"`
	MobilePhone             string                   `json:"mobile_phone"`
	MobileContactPermission bool                     `json:"mobile_contact_permission"`
	NewsUpdatesPermission   bool                     `json:"news_updates_permission"`
	Active                  bool                     `json:"active"`
	Paused                  bool                     `json:"paused"`
	PausedEndDate           string                   `json:"paused_end_date"`
	RecipientInfo           LasagnaLoveRecipientInfo `json:"recipient_info"`
	VolunteerInfo           LasagnaLoveVolunteerInfo `json:"volunteer_info"`
}
