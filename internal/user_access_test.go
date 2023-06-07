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
			call:    func() (model.LasagnaLoveUser, error) { return AuthorizeUser("TestUser1", "password1") },
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

func TestGetUserByUserName(t *testing.T) {
	tests := []userTestType{
		{
			name:    "Existing username",
			call:    func() (model.LasagnaLoveUser, error) { return GetUserByUserName("TestUser1") },
			wantErr: false,
		},
		{
			name:    "Non-existing username",
			call:    func() (model.LasagnaLoveUser, error) { return GetUserByUserName("NonExistingUser") },
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
					Username:   "TestUser3",
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
					Username:        "TestUser-invalid role",
					Password:        "password3",
					Email:           "testuser3@example.com",
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
			name: "Add duplicate user",
			call: func() (model.LasagnaLoveUser, error) {
				return AddNewUser(model.LasagnaLoveUser{
					Username:   "TestUser1",
					Password:   "password1",
					GivenName:  "Test",
					FamilyName: "UserOne",
				})
			},
			wantErr: true,
		},
		{
			name: "Add new user",
			call: func() (model.LasagnaLoveUser, error) {
				return AddNewUser(model.LasagnaLoveUser{
					Roles:           []string{"chef"},
					Username:        "TestUser3",
					Password:        "password3",
					Email:           "testuser3@example.com",
					GivenName:       "Test",
					FamilyName:      "UserThree",
					StreetAddress:   []string{"111 Testing Plaza", "Suite 1"},
					City:            "Anywhere",
					StateOrProvince: "AB",
					PostalCode:      "T5B 6W2",
					MobilePhone:     "780-555-1212",
				})
			},
			wantErr: false,
		},
	}
	runUserTests(t, tests)
}
