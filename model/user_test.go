package model_test

import (
	"github.com/bythecover/backend/model"
	"testing"
)

// create a custom struct

type testUser struct {
	Id        string
	FirstName string
	LastName  string
	Email     string
	Role      string
}

func TestBadUsers(t *testing.T) {
	testCases := []struct {
		input testUser
		want  string
	}{
		{
			input: testUser{
				Id:        "123",
				FirstName: "",
				LastName:  "Doe",
				Email:     "doe@email.com",
				Role:      "User",
			},
			want: model.ErrEmptyName.Error(),
		},
		{
			input: testUser{
				Id:        "123",
				FirstName: "John",
				LastName:  "",
				Email:     "doe@email.com",
				Role:      "User",
			},
			want: model.ErrEmptyName.Error(),
		},
		{
			input: testUser{
				Id:        "123",
				FirstName: "John",
				LastName:  "Doe",
				Email:     "",
				Role:      "User",
			},
			want: model.ErrEmptyEmail.Error(),
		},
		{
			input: testUser{
				Id:        "123",
				FirstName: "",
				LastName:  "",
				Email:     "",
				Role:      "User",
			},
			want: model.ErrEmptyName.Error(),
		},
	}

	for _, testCase := range testCases {
		testName := testCase.want
		t.Run(testName, func(t *testing.T) {
			_, err := model.NewUser(testCase.input.Id, testCase.input.FirstName, testCase.input.LastName, testCase.input.Email, testCase.input.Role)

			if err.Error() != testCase.want {
				t.Errorf("got %v, want %v", err, testCase.want)
			}
		})
	}
}
