package internal

import (
	"project-ricotta/bechamel-api/model"
)

// Project Ricotta: Bechamel API
//
// This is a temporary data source with dummy data.
// This is here to allow the Bechamel API portion of Project Ricotta to get started.
// This will be replaced with calls to the Ragu user information service,
// once that is available.

var LasagnaLoveUsersDummyData = []model.LasagnaLoveUser{
	{
		ID:                      1,
		Roles:                   []string{"chef"},
		Username:                "TestUser1",
		Password:                "EsX3b/B4fCYGb2+iAs4fAIXQtiq3EydUDi03ECVvTEE=", // "password1, hashed"
		GivenName:               "Test",
		FamilyName:              "UserOne",
		Email:                   "testuser1@example.com",
		EmailValidated:          true,
		CreationTime:            "2023-04-11T07:11:04.332Z",
		LastUpdateTime:          "2023-06-06T13:00:00.000Z",
		StreetAddress:           []string{"123 Fake St.", "Ground floor"},
		City:                    "Springfield",
		StateOrProvince:         "OH",
		PostalCode:              "45502",
		MobilePhone:             "937-555-1212",
		MobileContactPermission: true,
		NewsUpdatesPermission:   false,
		Active:                  true,
		Paused:                  false,
		Attestations: model.LasagnaLoveAttestations{
			UserIsEighteen:                   true,
			UserAcceptedEmailCommunications:  true,
			VolunteerAcceptedIndemnityWaiver: "2023-04-11T07:11:05.122Z",
			VolunteerAcceptedVolunteerTerms:  "2023-04-11T07:11:05.191Z",
			VolunteerCompletedSafetyTraining: "2023-04-11T08:41:19.706Z",
		},
		VolunteerInfo: &model.LasagnaLoveVolunteerInfo{
			Birthday:                 "1980-02-29",
			GenderIdentity:           "male",
			MaxTravelDistance:        10,
			FamiliesPerDelivery:      1,
			AccomodatesExtraRequests: false,
		},
	},
	{
		ID:                      2,
		Roles:                   []string{"requester", "recipient"},
		Username:                "TestUser2",
		Password:                "TnhbYUymFq5gr1jvyw1AmTviqlp3sYp7t0VxfT7ut1M=", // "password2", hashed
		GivenName:               "Test",
		FamilyName:              "UserTwo",
		Email:                   "testuser2@example.com",
		EmailValidated:          true,
		CreationTime:            "2023-05-16T06:44:19.794Z",
		LastUpdateTime:          "2023-06-06T13:00:00.000Z",
		StreetAddress:           []string{"999 False Way"},
		City:                    "Springfield",
		StateOrProvince:         "OH",
		PostalCode:              "45501",
		MobilePhone:             "937-555-8888",
		MobileContactPermission: true,
		NewsUpdatesPermission:   true,
		Active:                  true,
		Paused:                  false,
		Attestations: model.LasagnaLoveAttestations{
			UserIsEighteen:                    true,
			UserAcceptedEmailCommunications:   true,
			RequesterAcceptedLiabilityRelease: "2023-05-16T06:44:19.901Z",
		},
		RecipientInfo: &model.LasagnaLoveRecipientInfo{
			AdultCount:          4,
			ChildCount:          1,
			DietaryRestrictions: []string{"vegetarian"},
		},
	},
}

var LasagnaLoveRequests_DummyData = []model.LasagnaLoveRequest{
	{
		ID:             1,
		RequesterID:    2,
		RecipientID:    2,
		Type:           "meal",
		Stage:          "accepted",
		CreationTime:   "2023-06-01T23:01:39.211Z",
		LastUpdateTime: "2023-06-04T08:31:15.013Z",
		Notes:          "Dummy internal testing request for Project Ricotta project.",
	},
}
