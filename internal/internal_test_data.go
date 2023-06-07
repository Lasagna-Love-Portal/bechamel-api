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

var LasagnaLoveUsers_DummyData = []model.LasagnaLoveUser{
	{
		ID:                      1,
		Roles:                   []string{"chef"},
		Username:                "TestUser1",
		Password:                "password1",
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
		Password:                "password2",
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
		RecipientInfo: &model.LasagnaLoveRecipientInfo{
			AdultCount:          4,
			ChildCount:          1,
			DietaryRestrictions: []string{"vegetarian"},
		},
	},
}
