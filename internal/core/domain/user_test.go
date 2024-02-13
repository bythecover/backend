package domain_test

import (
	"bythecover/backend/internal/core/domain"
	"testing"
)

// create a custom struct

type testUser struct {
	FirstName string
	LastName  string
	Email     string
	IsAuthor  bool
}

func TestBadUsers(t *testing.T) {
	testCases := []struct {
		input testUser
		want  string
	}{
		{
			input: testUser{
				FirstName: "",
				LastName:  "Doe",
				Email:     "doe@email.com",
				IsAuthor:  false,
			},
			want: domain.ErrEmptyName.Error(),
		},
		{
			input: testUser{
				FirstName: "John",
				LastName:  "",
				Email:     "doe@email.com",
				IsAuthor:  false,
			},
			want: domain.ErrEmptyName.Error(),
		},
		{
			input: testUser{
				FirstName: "John",
				LastName:  "Doe",
				Email:     "",
				IsAuthor:  false,
			},
			want: domain.ErrEmptyEmail.Error(),
		},
		{
			input: testUser{
				FirstName: "",
				LastName:  "",
				Email:     "",
				IsAuthor:  false,
			},
			want: domain.ErrEmptyName.Error(),
		},
	}

	for _, testCase := range testCases {
		testName := testCase.want
		t.Run(testName, func(t *testing.T) {
			_, err := domain.NewUser(testCase.input.FirstName, testCase.input.LastName, testCase.input.Email, testCase.input.IsAuthor)

			if err.Error() != testCase.want {
				t.Errorf("got %v, want %v", err, testCase.want)
			}
		})
	}
}
