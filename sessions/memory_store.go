package sessions

import (
	"fmt"
	"log"

	"github.com/google/uuid"
)

// Memory Store is an in-memory implementation of SessionStore meant to
// be used for testing purposes
type MemoryStore map[string]*Session

func (store MemoryStore) Save(session *Session) uuid.UUID {
	sessionExists := session.Id.ID() != 0
	if !sessionExists {
		id := uuid.New()
		session.Id = id
	}
	store[session.Id.String()] = session
	return session.Id
}

func (store MemoryStore) Get(id string) (*Session, error) {
	if store[id] == nil {
		log.Println(fmt.Errorf("%w For id: %v", ErrSessionNotFound, id))
		return nil, ErrSessionNotFound
	}
	return store[id], nil
}
