package internal

import (
	"project-ricotta/bechamel-api/model"
	"testing"
)

func TestAuthorizeUser(t *testing.T) {
	tests := []struct {
		name     string
		userName string
		password string
		wantErr  bool
	}{
		{
			name:     "Correct credentials",
			userName: "TestUser1",
			password: "password1",
			wantErr:  false,
		},
		{
			name:     "Empty credentials",
			userName: "",
			password: "",
			wantErr:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := AuthorizeUser(tt.userName, tt.password)
			if (err != nil) != tt.wantErr {
				t.Errorf("AuthorizeUser() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestGetUserByID(t *testing.T) {
	tests := []struct {
		name    string
		userID  int
		wantErr bool
	}{
		{
			name:    "Existing user ID",
			userID:  1,
			wantErr: false,
		},
		{
			name:    "Non-existing user ID",
			userID:  100,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := GetUserByID(tt.userID)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetUserByID() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestGetUserByUserName(t *testing.T) {
	tests := []struct {
		name     string
		userName string
		wantErr  bool
	}{
		{
			name:     "Existing username",
			userName: "TestUser1",
			wantErr:  false,
		},
		{
			name:     "Non-existing username",
			userName: "NonExistingUser",
			wantErr:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := GetUserByUserName(tt.userName)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetUserByUserName() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestAddNewUser(t *testing.T) {
	tests := []struct {
		name           string
		newUserProfile model.LasagnaLoveUser
		wantErr        bool
	}{
		{
			name: "Add valid new user",
			newUserProfile: model.LasagnaLoveUser{
				Username:   "TestUser3",
				Password:   "password3",
				GivenName:  "Test",
				FamilyName: "UserThree",
			},
			wantErr: false,
		},
		{
			name: "Add duplicate user",
			newUserProfile: model.LasagnaLoveUser{
				Username:   "TestUser1",
				Password:   "password1",
				GivenName:  "Test",
				FamilyName: "UserOne",
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := AddNewUser(tt.newUserProfile)
			if (err != nil) != tt.wantErr {
				t.Errorf("AddNewUser() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
