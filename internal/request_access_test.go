package internal

import (
	"project-ricotta/bechamel-api/model"
	"testing"
)

type requestTestType struct {
	name    string
	call    func() (model.LasagnaLoveRequest, error)
	wantErr bool
}

func runRequestTests(t *testing.T, tests []requestTestType) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := tt.call()
			if (err != nil) != tt.wantErr {
				t.Errorf("error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestGetRequestByID(t *testing.T) {
	tests := []requestTestType{
		{
			name:    "Existing request ID",
			call:    func() (model.LasagnaLoveRequest, error) { return GetRequestByID(1) },
			wantErr: false,
		},
		{
			name:    "Invalid request ID",
			call:    func() (model.LasagnaLoveRequest, error) { return GetRequestByID(0) },
			wantErr: true,
		},
	}
	runRequestTests(t, tests)
}

func TestAddNewRequest(t *testing.T) {
	tests := []requestTestType{
		{
			name: "Add new request with insufficient data",
			call: func() (model.LasagnaLoveRequest, error) {
				return AddNewRequest(model.LasagnaLoveRequest{
					Stage: "contacted",
				})
			},
			wantErr: true,
		},
		{
			name: "Add new request with ID specified",
			call: func() (model.LasagnaLoveRequest, error) {
				return AddNewRequest(model.LasagnaLoveRequest{
					ID:    33,
					Stage: "contacted",
				})
			},
			wantErr: true,
		},
		{
			name: "Add new request with invalid stage value",
			call: func() (model.LasagnaLoveRequest, error) {
				return AddNewRequest(model.LasagnaLoveRequest{
					RequesterID: 1,
					RecipientID: 1,
					Stage:       "garbage stage",
				})
			},
			wantErr: true,
		},
		{
			name: "Add new request with invalid type value",
			call: func() (model.LasagnaLoveRequest, error) {
				return AddNewRequest(model.LasagnaLoveRequest{
					RequesterID: 1,
					RecipientID: 1,
					Type:        "horse stall mucking request",
				})
			},
			wantErr: true,
		},
		{
			name: "Add new request",
			call: func() (model.LasagnaLoveRequest, error) {
				return AddNewRequest(model.LasagnaLoveRequest{
					RequesterID: 1,
					RecipientID: 1,
				})
			},
			wantErr: false,
		},
	}
	runRequestTests(t, tests)
}
