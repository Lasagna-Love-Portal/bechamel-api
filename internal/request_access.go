package internal

// Project Ricotta: Bechamel API
//
// TODO: this is coupled somewhat tightly to the dummy test data held
// in internal_test_data.go - as we add external data source access,
// this will need to change to match.

import (
	"errors"
	"project-ricotta/bechamel-api/model"
	"time"
)

var permittedStages = [...]string{"ingested", "reviewed", "accepted",
	"backlog", "matched", "contacted", "scheduled",
	"delivered", "no_response", "no_longer_wanted"}

func findRequest(requestFilter func(model.LasagnaLoveRequest) bool) (model.LasagnaLoveRequest, error) {
	for _, request := range LasagnaLoveRequests_DummyData {
		if requestFilter(request) {
			return request, nil
		}
	}
	return model.LasagnaLoveRequest{}, errors.New("no user found with the supplied criteria")
}

func GetRequestByID(requestID int) (model.LasagnaLoveRequest, error) {
	return findRequest(func(r model.LasagnaLoveRequest) bool { return r.ID == requestID })
}

func AddNewRequest(newRequest model.LasagnaLoveRequest) (model.LasagnaLoveRequest, error) {
	// Not allowed to specify an ID for a request - error if one is provided
	if newRequest.ID != 0 {
		return model.LasagnaLoveRequest{}, errors.New("request ID may not be specified")
	} else if newRequest.RequesterID == 0 {
		return model.LasagnaLoveRequest{}, errors.New("request missing required RequesterID value")
	} else if newRequest.RecipientID == 0 {
		return model.LasagnaLoveRequest{}, errors.New("request missing required RecipientID value")
	} else if newRequest.Stage != "" && !StringIsInArray(permittedStages[:], newRequest.Stage) {
		return model.LasagnaLoveRequest{}, errors.New("request stage not a permitted value")
	} else if newRequest.Type != "" && newRequest.Type != "meal" {
		return model.LasagnaLoveRequest{}, errors.New("request type must be \"meal\" if specified")
	}

	newRequest.ID = len(LasagnaLoveRequests_DummyData) + 1
	if newRequest.Stage == "" {
		newRequest.Stage = "ingested"
	}
	if newRequest.Type == "" {
		newRequest.Type = "meal"
	}
	// NOTE: this is not an arbitrary formatting string, this is required format string
	// to get time created in ISO 8601 simplified extended format as returned
	// by JavaScript's toISOString() function.
	newRequest.CreationTime = time.Now().UTC().Format("2006-01-02T15:04:05.000Z0700")
	newRequest.LastUpdateTime = newRequest.CreationTime

	LasagnaLoveRequests_DummyData = append(LasagnaLoveRequests_DummyData, newRequest)
	return newRequest, nil
}
