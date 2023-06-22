package model

// Authorization requests should have either the username/password pair,
// or the refresh_token. If both sets of credentials are supplied,
// the refresh_token is used.
// TODO: Originally wanted to separate into two different models,
// one with just username/password and one with refresh_token; both with the fields
// having a binding:required entry.
// However, gin's BindJSON returns an EOF error if called twice
// with the same gin Context, which seems to prevent checking
// against two JSON schemas in succession
type LasagnaLoveAuthRequest struct {
	Email        string `json:"email"`
	Password     string `json:"password"`
	RefreshToken string `json:"refresh_token"`
}
