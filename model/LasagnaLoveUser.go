package model

import "encoding/json"

type LasagnaLoveUser struct {
	UserID     int    `json:"id"`
	UserName   string `json:"userName"`
	Password   string `json:"password"`
	GivenName  string `json:"givenName"`
	MiddleName string `json:"middleName"`
	FamilyName string `json:"familyName"`
}

/* This overrides the default marshaling of the structure to JSON, removing the password field value.
It has the unfortunate side effect of leaving an empty string entry in the generated JSON
which leaks out to users, but does allow the use of the same struct for GETting and POSTing user profiles.

Better ideas are welcomed.
*/
func (l LasagnaLoveUser) MarshalJSON() ([]byte, error) {
	type lasagnaLoveUser LasagnaLoveUser // prevent recursion
	x := lasagnaLoveUser(l)
	x.Password = ""
	return json.Marshal(x)
}
