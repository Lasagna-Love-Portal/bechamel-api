package model

type LasagnaLoveUser struct {
	ID                 int    `json:"id"`
	Username           string `json:"username"`
	Password           string `json:"-"`
	GivenName          string `json:"given_name"`
	MiddleOrMaidenName string `json:"middle_or_maiden_name"`
	FamilyName         string `json:"family_name"`
}
