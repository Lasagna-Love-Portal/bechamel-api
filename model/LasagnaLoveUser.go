package model

import "encoding/json"

type LasagnaLoveUser struct {
	ID                 int    `json:"id"`
	Username           string `json:"username"`
	Password           string `json:"password"`
	GivenName          string `json:"given_name"`
	MiddleOrMaidenName string `json:"middle_or_maiden_name"`
	FamilyName         string `json:"family_name"`
}

// This overrides the default marshaling of the structure to JSON, removing the password field value.
// It has the unfortunate side effect of leaving an empty string entry in the generated JSON
// which leaks out to users, but does allow the use of the same struct for GETting and POSTing user profiles.
//
// Better ideas are welcomed.
func (l LasagnaLoveUser) MarshalJSON() ([]byte, error) {
	type lasagnaLoveUser LasagnaLoveUser // prevent recursion
	x := lasagnaLoveUser(l)
	x.Password = ""
	return json.Marshal(x)
}
