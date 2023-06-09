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

func findUser(userFilter func(model.LasagnaLoveUser) bool) (model.LasagnaLoveUser, error) {
	for _, user := range LasagnaLoveUsersDummyData {
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

func GetUserByEmailAddress(emailAddress string) (model.LasagnaLoveUser, error) {
	if emailAddress == "" {
		return model.LasagnaLoveUser{}, errors.New("emailAddress must be non-empty")
	}

	return findUser(func(u model.LasagnaLoveUser) bool { return u.Email == emailAddress })
}

func AddNewUser(newUserProfile model.LasagnaLoveUser) (model.LasagnaLoveUser, error) {
	// Not allowed to specify an userID - error if one is provided
	if newUserProfile.ID != 0 {
		return model.LasagnaLoveUser{}, errors.New("userID may not be specified")
	}
	if len(newUserProfile.Roles) == 0 ||
		newUserProfile.Username == "" ||
		newUserProfile.Password == "" ||
		newUserProfile.Email == "" ||
		newUserProfile.GivenName == "" ||
		newUserProfile.FamilyName == "" ||
		len(newUserProfile.StreetAddress) == 0 ||
		newUserProfile.City == "" ||
		newUserProfile.StateOrProvince == "" ||
		newUserProfile.PostalCode == "" ||
		newUserProfile.MobilePhone == "" {
		return model.LasagnaLoveUser{}, errors.New("invalid or incomplete user data")
	}
	for _, role := range newUserProfile.Roles {
		if !StringIsInArray(model.LasagnaLoveUserPermittedRoles[:], role) {
			return model.LasagnaLoveUser{}, errors.New("invalid value supplied in roles array")
		}
	}
	if _, err := GetUserByUserName(newUserProfile.Username); err == nil {
		return model.LasagnaLoveUser{}, errors.New("username already exists")
	}
	if _, err := GetUserByEmailAddress(newUserProfile.Email); err == nil {
		return model.LasagnaLoveUser{}, errors.New("email address already in use, dupliate usage not permitted")
	}

	newUserProfile.ID = len(LasagnaLoveUsersDummyData) + 1
	newUserProfile.Password = HashPassword(newUserProfile.Password)
	// NOTE: this is not an arbitrary formatting string, this is required format string
	// to get time created in ISO 8601 simplified extended format as returned
	// by JavaScript's toISOString() function.
	newUserProfile.CreationTime = time.Now().UTC().Format("2006-01-02T15:04:05.000Z0700")
	newUserProfile.LastUpdateTime = newUserProfile.CreationTime
	LasagnaLoveUsersDummyData = append(LasagnaLoveUsersDummyData, newUserProfile)
	return newUserProfile, nil
}
