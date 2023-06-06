package internal

import (
	"project-ricotta/bechamel-api/model"
	"testing"
)

type testType struct {
	name    string
	call    func() (model.LasagnaLoveUser, error)
	wantErr bool
}

func runTests(t *testing.T, tests []testType) {
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
	tests := []testType{
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
	runTests(t, tests)
}

func TestGetUserByID(t *testing.T) {
	tests := []testType{
		{
			name:    "Existing user ID",
			call:    func() (model.LasagnaLoveUser, error) { return GetUserByID(1) },
			wantErr: false,
		},
		{
			name:    "Non-existing user ID",
			call:    func() (model.LasagnaLoveUser, error) { return GetUserByID(100) },
			wantErr: true,
		},
	}
	runTests(t, tests)
}

func TestGetUserByUserName(t *testing.T) {
	tests := []testType{
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
	runTests(t, tests)
}

func TestAddNewUser(t *testing.T) {
	tests := []testType{
		{
			name: "Add valid new user",
			call: func() (model.LasagnaLoveUser, error) {
				return AddNewUser(model.LasagnaLoveUser{
					Username:   "TestUser3",
					Password:   "password3",
					GivenName:  "Test",
					FamilyName: "UserThree",
				})
			},
			wantErr: false,
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
	}
	runTests(t, tests)
}
