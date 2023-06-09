package internal

// Project Ricotta: Bechamel API
//
// This is a temporary data source with dummy data.
// This is here to allow the Bechamel API portion of Project Ricotta to get started.
// This will be replaced with calls to the Ragu user information service,
// once that is available.

import (
	"errors"
	"project-ricotta/bechamel-api/model"
)

var lasagnaLoveUsers = []model.LasagnaLoveUser{
	{ID: 1, Username: "TestUser1", Password: "EsX3b/B4fCYGb2+iAs4fAIXQtiq3EydUDi03ECVvTEE=", // "password1"
		GivenName: "Test", FamilyName: "UserOne"},
	{ID: 2, Username: "TestUser2", Password: "TnhbYUymFq5gr1jvyw1AmTviqlp3sYp7t0VxfT7ut1M=", // "password2"
		GivenName: "Test", FamilyName: "UserTwo"},
}

func findUser(userFilter func(model.LasagnaLoveUser) bool) (model.LasagnaLoveUser, error) {
	for _, user := range lasagnaLoveUsers {
		if userFilter(user) {
			return user, nil
		}
	}
	return model.LasagnaLoveUser{}, errors.New("no user found with the supplied criteria")
}

func AuthorizeUser(userName string, password string) (model.LasagnaLoveUser, error) {
	if userName == "" || password == "" {
		return model.LasagnaLoveUser{}, errors.New("userName and password must be non-empty")
	}

	return findUser(func(u model.LasagnaLoveUser) bool {
		return u.Username == userName && u.Password == HashPassword(password)
	})
}

func GetUserByID(userID int) (model.LasagnaLoveUser, error) {
	return findUser(func(u model.LasagnaLoveUser) bool { return u.ID == userID })
}

func GetUserByUserName(userName string) (model.LasagnaLoveUser, error) {
	if userName == "" {
		return model.LasagnaLoveUser{}, errors.New("userName must be non-empty")
	}

	return findUser(func(u model.LasagnaLoveUser) bool { return u.Username == userName })
}

func AddNewUser(newUserProfile model.LasagnaLoveUser) (model.LasagnaLoveUser, error) {
	// Not allowed to specify an userID - error if one is provided
	if newUserProfile.ID != 0 {
		return model.LasagnaLoveUser{}, errors.New("userID may not be specified")
	}

	if newUserProfile.FamilyName == "" || newUserProfile.GivenName == "" || newUserProfile.Username == "" || newUserProfile.Password == "" {
		return model.LasagnaLoveUser{}, errors.New("invalid or incomplete user data")
	}

	if _, err := GetUserByUserName(newUserProfile.Username); err == nil {
		return model.LasagnaLoveUser{}, errors.New("username already exists")
	}

	newUserProfile.ID = len(lasagnaLoveUsers) + 1
	newUserProfile.Password = HashPassword(newUserProfile.Password)
	lasagnaLoveUsers = append(lasagnaLoveUsers, newUserProfile)
	return newUserProfile, nil
}
