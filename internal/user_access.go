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
	"strings"
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

/*
	TODO: not yet able to use this function in UpdateUser below, run-time assertion comes up.

	func isKeyInStruct(key string, structToCheck interface{}) bool {
		rfl := reflect.ValueOf(&structToCheck).Elem()
		return rfl.FieldByName(key).IsValid()
	}
*/

func UpdateUser(userProfile model.LasagnaLoveUser, updates map[string]interface{}) (model.LasagnaLoveUser, error) {
	var llVolunteerInfo model.LasagnaLoveVolunteerInfo
	var llRecipientInfo model.LasagnaLoveRecipientInfo
	var didMakeUpdates bool
	var userID = userProfile.ID

	// The double loop here is intentional - this is to prevent partial updates
	// by making sure all fields supplied are valid before making any updates
	for key, value := range updates {

		rfl := reflect.ValueOf(&userProfile).Elem()
		if fld := rfl.FieldByName(key); !fld.IsValid() {
			return model.LasagnaLoveUser{}, errors.New("invalid field name supplied for update")
		}

		/*
			if !isKeyInStruct(key, userProfile) {
				return model.LasagnaLoveUser{}, errors.New("invalid field name supplied for update")
			}*/
		switch key {
		// Fields that are valid but not permitted to be updated
		case "Id": // this is a bit weird, is being picked up above as invalid field name
			fallthrough
		case "CreationTime":
			fallthrough
		case "LastUpdateTime":
			return model.LasagnaLoveUser{}, errors.New("updates contain field name that is not permitted to be updated")
		case "Attestations":
			for attKey := range value.(model.PatchUpdateStruct) {
				/*
					if !isKeyInStruct(attKey, userProfile.Attestations) {
						return model.LasagnaLoveUser{}, errors.New("invalid field name supplied for attestations update")
					}*/
				attRfl := reflect.ValueOf(&(userProfile.Attestations)).Elem()
				if attFld := attRfl.FieldByName(attKey); !attFld.IsValid() {
					return model.LasagnaLoveUser{}, errors.New("invalid field name supplied for attestations update")
				}
			}
		case "RecipientInfo":
			for recKey := range value.(model.PatchUpdateStruct) {
				recRfl := reflect.ValueOf(&llRecipientInfo).Elem()
				if recFld := recRfl.FieldByName(recKey); !recFld.IsValid() {
					return model.LasagnaLoveUser{}, errors.New("invalid field name supplied for recipient_info update")
				}
				// TODO: make sure value types for incoming updates are compatible
			}
		case "VolunteerInfo":
			for volKey := range value.(model.PatchUpdateStruct) {
				volRfl := reflect.ValueOf(&llVolunteerInfo).Elem()
				if volFld := volRfl.FieldByName(volKey); !volFld.IsValid() {
					return model.LasagnaLoveUser{}, errors.New("invalid field name supplied for volunteer_info update")
				}
				// TODO: make sure value types for incoming updates are compatible
			}
		}
	}

	// TODO: adding and updating will likely need to be datastore dependent and not common.
	// For the integrated fixed data, note the switch to directly accessing the LasagnaLoveUsersDummyData here.
	for key, value := range updates {
		switch key {
		case "Password":
			LasagnaLoveUsersDummyData[userID-1].Password = HashPassword(reflect.ValueOf(value).String())
		case "RecipientInfo":
			if LasagnaLoveUsersDummyData[userID-1].RecipientInfo == nil {
				LasagnaLoveUsersDummyData[userID-1].RecipientInfo = &llRecipientInfo
			}
			for recKey, recValue := range value.(model.PatchUpdateStruct) {
				if strings.HasPrefix(reflect.TypeOf(recValue).String(), "float") {
					reflect.ValueOf(LasagnaLoveUsersDummyData[userID-1].RecipientInfo).Elem().FieldByName(recKey).SetInt(
						int64(reflect.ValueOf(recValue).Float()))
				} else {
					reflect.ValueOf(LasagnaLoveUsersDummyData[userID-1].RecipientInfo).Elem().FieldByName(recKey).Set(
						reflect.ValueOf(recValue))
				}
			}
		case "VolunteerInfo":
			if LasagnaLoveUsersDummyData[userID-1].VolunteerInfo == nil {
				LasagnaLoveUsersDummyData[userID-1].VolunteerInfo = &llVolunteerInfo
			}
			for volKey, volValue := range value.(model.PatchUpdateStruct) {
				if strings.HasPrefix(reflect.TypeOf(volValue).String(), "float") {
					reflect.ValueOf(LasagnaLoveUsersDummyData[userID-1].VolunteerInfo).Elem().FieldByName(volKey).SetInt(
						int64(reflect.ValueOf(volValue).Float()))
				} else {
					reflect.ValueOf(LasagnaLoveUsersDummyData[userID-1].VolunteerInfo).Elem().FieldByName(volKey).Set(
						reflect.ValueOf(volValue))
				}
			}
		case "Attestations":
			for attKey, attValue := range value.(model.PatchUpdateStruct) {
				reflect.ValueOf(&LasagnaLoveUsersDummyData[userID-1].Attestations).Elem().FieldByName(attKey).Set(
					reflect.ValueOf(attValue))
			}
		default:
			reflect.ValueOf(&LasagnaLoveUsersDummyData[userID-1]).Elem().FieldByName(key).Set(reflect.ValueOf(value))
		}
		didMakeUpdates = true
	}
	if didMakeUpdates {
		LasagnaLoveUsersDummyData[userID-1].LastUpdateTime = CurrentTimeAsISO8601String()
	}

	return LasagnaLoveUsersDummyData[userID-1], nil
}
