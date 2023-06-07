package model

import "encoding/json"

// TODO: Fix up required vs. optional
type LasagnaLoveUser struct {
	ID                      int                      `json:"id"`
	Roles                   []string                 `json:"roles"`
	Username                string                   `json:"username"`
	Password                string                   `json:"password"`
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

// This overrides the default marshaling of the structure to JSON, removing the password field value.
// It has the unfortunate side effect of leaving an empty string entry in the generated JSON
// which leaks out to users, but does allow the use of the same struct for GETting and POSTing user profiles.
//
// Simply using a marshalling entry of "-" in the struct above causes the marshalling in the POST calls
// to create user Profiles to not unmarshall the supplied password field.
//
// TODO: Better ideas to accomplish allowing and reading the "password" JSON field value when creating
// LasagnaLoveUser data structures, while not emitting these when marshalling, are welcomed.
func (l LasagnaLoveUser) MarshalJSON() ([]byte, error) {
	type lasagnaLoveUser LasagnaLoveUser // prevent recursion
	x := lasagnaLoveUser(l)
	x.Password = ""
	return json.Marshal(x)
}
