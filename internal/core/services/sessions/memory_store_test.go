package sessions_test

import (
	"bythecover/backend/internal/core/services/sessions"
	"errors"
	"testing"
)

type output struct {
	session *sessions.Session
	err     error
}

type errorTestCase struct {
	input  string
	output output
	name   string
}

func TestMemoryStoreGet(t *testing.T) {
	store := make(sessions.MemoryStore)
	validSession := &sessions.Session{}
	store["321"] = validSession

	testTable := []errorTestCase{
		{
			input: "123",
			output: output{
				session: nil,
				err:     sessions.ErrSessionNotFound,
			},
			name: "not found session id",
		},
		{
			input: "321",
			output: output{
				session: validSession,
				err:     nil,
			},
			name: "Valid Get",
		},
	}

	for _, scenario := range testTable {
		t.Run(scenario.name, func(t *testing.T) {
			session, err := store.Get(scenario.input)
			if !errors.Is(err, scenario.output.err) {
				t.Fatalf("Wanted: '%v', got '%v'", scenario.output.err, err)
			}
			if session != scenario.output.session {
				t.Fatalf("Session mistmatch Wanted: '%v', got '%v'", scenario.output.session, session)
			}
		})
	}

}
