package model

// TODO: should the model be using IDs for the requester and recipient?
// Or should the model be using pointers-to-info blocks?
type LasagnaLoveRequest struct {
	ID             int    `json:"id"`
	RequesterID    int    `json:"requester_id"`
	RecipientID    int    `json:"recipient_id"`
	Type           string `json:"type"`
	Stage          string `json:"stage"`
	CreationTime   string `json:"creation_time"`
	LastUpdateTime string `json:"last_update_time"`
	Notes          string `json:"notes"`
}

var LasagnaLoveRequestPermittedStages = [...]string{"ingested", "reviewed", "accepted",
	"backlog", "matched", "contacted", "scheduled",
	"delivered", "no_response", "no_longer_wanted"}
