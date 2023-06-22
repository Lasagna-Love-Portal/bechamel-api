package internal

import (
	"project-ricotta/bechamel-api/model"
	"testing"
)

type userTestType struct {
	name    string
	call    func() (model.LasagnaLoveUser, error)
	wantErr bool
}

func runUserTests(t *testing.T, tests []userTestType) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := tt.call()
			if (err != nil) != tt.wantErr {
				t.Errorf("error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestAuthorizeUser(t *testing.T) {
	tests := []userTestType{
		{
			name:    "Correct credentials",
			call:    func() (model.LasagnaLoveUser, error) { return AuthorizeUser("testuser1@example.com", "password1") },
			wantErr: false,
		},
		{
			name:    "Empty credentials",
			call:    func() (model.LasagnaLoveUser, error) { return AuthorizeUser("", "") },
			wantErr: true,
		},
	}
	runUserTests(t, tests)
}

func TestGetUserByID(t *testing.T) {
	tests := []userTestType{
		{
			name:    "Existing user ID",
			call:    func() (model.LasagnaLoveUser, error) { return GetUserByID(1) },
			wantErr: false,
		},
		{
			name:    "Non-existing user ID",
			call:    func() (model.LasagnaLoveUser, error) { return GetUserByID(0) },
			wantErr: true,
		},
	}
	runUserTests(t, tests)
}

func TestAddNewUser(t *testing.T) {
	tests := []userTestType{
		{
			name: "Add new user with insufficient data",
			call: func() (model.LasagnaLoveUser, error) {
				return AddNewUser(model.LasagnaLoveUser{
					Password:   "password3",
					GivenName:  "Test",
					FamilyName: "UserThree",
				})
			},
			wantErr: true,
		},
		{
			name: "Add new user with invalid role in Roles array",
			call: func() (model.LasagnaLoveUser, error) {
				return AddNewUser(model.LasagnaLoveUser{
					Roles:           []string{"chef", "invalid"},
					Email:           "testuser3@example.com",
					Password:        "password3",
					GivenName:       "Test",
					FamilyName:      "UserThree",
					StreetAddress:   []string{"111 Testing Plaza", "Suite 1"},
					City:            "Anywhere",
					StateOrProvince: "AB",
					PostalCode:      "T5B 6W2",
					MobilePhone:     "780-555-1212"})
			},
			wantErr: true,
		},
		{
			name: "Add new user",
			call: func() (model.LasagnaLoveUser, error) {
				return AddNewUser(model.LasagnaLoveUser{
					Roles:           []string{"chef"},
					Email:           "testuser3@example.com",
					Password:        "password3",
					GivenName:       "Test",
					FamilyName:      "UserThree",
					StreetAddress:   []string{"111 Testing Plaza", "Suite 1"},
					City:            "Anywhere",
					StateOrProvince: "AB",
					Country:         "US",
					PostalCode:      "T5B 6W2",
					MobilePhone:     "780-555-1212",
				})
			},
			wantErr: false,
		},
		{
			name: "Add user with duplicated Email address",
			call: func() (model.LasagnaLoveUser, error) {
				return AddNewUser(model.LasagnaLoveUser{
					Roles:           []string{"chef"},
					Email:           "testuser3@example.com",
					Password:        "password3",
					GivenName:       "Test",
					FamilyName:      "UserThree",
					StreetAddress:   []string{"111 Testing Plaza", "Suite 1"},
					City:            "Anywhere",
					StateOrProvince: "AB",
					PostalCode:      "T5B 6W2",
					MobilePhone:     "780-555-1212",
				})
			},
			wantErr: true,
		},
	}
	runUserTests(t, tests)
}
