package internal

// Project Ricotta: Bechamel API
//
// TODO: this is coupled somewhat tightly to the dummy test data held
// in internal_test_data.go - as we add external data source access,
// this will need to change to match.

import (
	"errors"
	"project-ricotta/bechamel-api/model"
	"reflect"
)

func findUser(userFilter func(model.LasagnaLoveUser) bool) (model.LasagnaLoveUser, error) {
	for _, user := range LasagnaLoveUsersDummyData {
		if userFilter(user) {
			return user, nil
		}
	}
	return model.LasagnaLoveUser{}, errors.New("no user found with the supplied criteria")
}

func AuthorizeUser(emailAddress string, password string) (model.LasagnaLoveUser, error) {
	if emailAddress == "" || password == "" {
		return model.LasagnaLoveUser{}, errors.New("email and password must be non-empty")
	}

	return findUser(func(u model.LasagnaLoveUser) bool {
		return u.Email == emailAddress && u.Password == HashPassword(password)
	})
}

func GetUserByID(userID int) (model.LasagnaLoveUser, error) {
	return findUser(func(u model.LasagnaLoveUser) bool { return u.ID == userID })
}

func GetUserByEmailAddress(emailAddress string) (model.LasagnaLoveUser, error) {
	if emailAddress == "" {
		return model.LasagnaLoveUser{}, errors.New("email must be non-empty")
	}

	return findUser(func(u model.LasagnaLoveUser) bool { return u.Email == emailAddress })
}

func AddNewUser(newUserProfile model.LasagnaLoveUser) (model.LasagnaLoveUser, error) {
	// Not allowed to specify an userID - error if one is provided
	if newUserProfile.ID != 0 {
		return model.LasagnaLoveUser{}, errors.New("profile ID may not be specified")
	}
	if len(newUserProfile.Roles) == 0 ||
		newUserProfile.Password == "" ||
		newUserProfile.Email == "" ||
		newUserProfile.GivenName == "" ||
		newUserProfile.FamilyName == "" ||
		len(newUserProfile.StreetAddress) == 0 ||
		newUserProfile.City == "" ||
		newUserProfile.StateOrProvince == "" ||
		newUserProfile.Country == "" ||
		newUserProfile.PostalCode == "" ||
		newUserProfile.MobilePhone == "" {
		return model.LasagnaLoveUser{}, errors.New("invalid or incomplete user data")
	}
	for _, role := range newUserProfile.Roles {
		if !StringIsInArray(model.LasagnaLoveUserPermittedRoles[:], role) {
			return model.LasagnaLoveUser{}, errors.New("invalid value supplied in roles array")
		}
	}
	if _, err := GetUserByEmailAddress(newUserProfile.Email); err == nil {
		return model.LasagnaLoveUser{}, errors.New("email address already in use, dupliate usage not permitted")
	}

	newUserProfile.ID = len(LasagnaLoveUsersDummyData) + 1
	newUserProfile.Password = HashPassword(newUserProfile.Password)
	newUserProfile.CreationTime = CurrentTimeAsISO8601String()
	newUserProfile.LastUpdateTime = newUserProfile.CreationTime
	LasagnaLoveUsersDummyData = append(LasagnaLoveUsersDummyData, newUserProfile)
	return newUserProfile, nil
}

func UpdateUser(userID int, updates map[string]any) (model.LasagnaLoveUser, error) {
	didMakeUpdates := false

	var currentUserProfile model.LasagnaLoveUser
	_, err := GetUserByID(userID)
	if err != nil {
		return model.LasagnaLoveUser{}, errors.New("could not obtain user for supplied ID")
	}

	// The double loop here is intentional - this is to prevent partial updates
	// by making sure all fields supplied are valid before making any updates
	for key := range updates {
		rfl := reflect.ValueOf(&currentUserProfile).Elem()
		if fld := rfl.FieldByName(key); !fld.IsValid() {
			return model.LasagnaLoveUser{}, errors.New("invalid field name supplied for update")
		}
		// TODO: unaccepted field handling
		// TODO: nested types/type ptrs (e.g. attestations)
	}

	// TODO: adding and updating will likely need to be datastore dependent and not common.
	// For the integrated fixed data, note the switch to directly accessing the LasagnaLoveUsersDummyData here.
	// TODO: need to set the referenced data structures as well, otherwise they need to be filled in full?
	for key, value := range updates {
		// TODO: password handling
		// TODO: nested types/type ptrs (e.g. attestations)
		reflect.ValueOf(&LasagnaLoveUsersDummyData[userID-1]).Elem().FieldByName(key).Set(reflect.ValueOf(value))
		didMakeUpdates = true
	}
	if didMakeUpdates {
		LasagnaLoveUsersDummyData[userID-1].LastUpdateTime = CurrentTimeAsISO8601String()
	}

	return LasagnaLoveUsersDummyData[userID-1], nil
}
