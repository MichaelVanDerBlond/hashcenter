package session

import (
	"time"

	"github.com/google/uuid"
)

func New() *Session {

	return &Session{
		ID:      uuid.NewString(),
		State:   Created,
		Started: time.Now(),
	}
}
