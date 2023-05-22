package model

import "encoding/json"

type LasagnaLoveUser struct {
	ID         int    `json:"ID"`
	Username   string `json:"Username"`
	Password   string `json:"Password"`
	GivenName  string `json:"given_Name"`
	MiddleName string `json:"middle_name"`
	FamilyName string `json:"family_name"`
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
