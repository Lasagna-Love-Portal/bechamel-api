package internal

/*
   Project Ricotta: Bechamel API

   This is a temporary data source with dummy data.
   This is here to allow the Bechamel API portion of Project Ricotta to get started.

   This will be replaced with calls to the Ragu user information service,
   once that is available.
*/
import (
	"errors"
	"project-ricotta/bechamel-api/model"

	"golang.org/x/exp/slices"
)

var lasagnaLoveUsers = []model.LasagnaLoveUser{
	{UserID: 1, UserName: "TestUser1", Password: "Password1", GivenName: "Test", FamilyName: "UserOne"},
	{UserID: 2, UserName: "TestUser2", Password: "Password2", GivenName: "Test", FamilyName: "UserTwo"},
}

func AuthorizeUser(userName string, password string) (model.LasagnaLoveUser, error) {
	if userName == "" || password == "" {
		return model.LasagnaLoveUser{}, errors.New("userName and password must be non-empty")
	}

	idx := slices.IndexFunc(lasagnaLoveUsers,
		func(l model.LasagnaLoveUser) bool { return l.UserName == userName && l.Password == password })
	if idx == -1 {
		return model.LasagnaLoveUser{}, errors.New("no user with supplied userName and password found")
	}
	return lasagnaLoveUsers[idx], nil
}

func GetUserByID(userID int) (model.LasagnaLoveUser, error) {
	idx := slices.IndexFunc(lasagnaLoveUsers,
		func(l model.LasagnaLoveUser) bool { return l.UserID == userID })
	if idx == -1 {
		return model.LasagnaLoveUser{}, errors.New("no user with supplied userID found")
	}
	return lasagnaLoveUsers[idx], nil
}

func GetUserByUserName(userName string) (model.LasagnaLoveUser, error) {
	if userName == "" {
		return model.LasagnaLoveUser{}, errors.New("userName must be non-empty")
	}
	idx := slices.IndexFunc(lasagnaLoveUsers,
		func(l model.LasagnaLoveUser) bool { return l.UserName == userName })
	if idx == -1 {
		return model.LasagnaLoveUser{}, errors.New("no user with supplied userName found")
	}
	return lasagnaLoveUsers[idx], nil
}

func AddNewUser(newUserProfile model.LasagnaLoveUser) (model.LasagnaLoveUser, error) {
	// Not allowed to specify an userID - error if one is provided
	if newUserProfile.UserID != 0 {
		return model.LasagnaLoveUser{}, errors.New("userID may not be specified")
	}

	// Verify required fields are present. Probably an easier way to do this...
	if newUserProfile.FamilyName == "" || newUserProfile.GivenName == "" ||
		newUserProfile.UserName == "" || newUserProfile.Password == "" {
		return model.LasagnaLoveUser{}, errors.New("one or more required fields missing or empty")
	}

	newUserProfile.UserID = len(lasagnaLoveUsers) + 1
	lasagnaLoveUsers = append(lasagnaLoveUsers, newUserProfile)
	return newUserProfile, nil
}
