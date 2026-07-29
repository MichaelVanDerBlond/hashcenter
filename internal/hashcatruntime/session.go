package hashcatruntime

import (
	"sync"
	"time"
)

type Session struct {
	mu sync.RWMutex

	snapshot Snapshot
}

func NewSession(name string) *Session {

	now := time.Now()

	s := &Session{}

	s.snapshot = Snapshot{
		Session: name,
		Started: now,
		Updated: now,
	}

	return s
}

func (s *Session) Snapshot() Snapshot {

	s.mu.RLock()
	defer s.mu.RUnlock()

	return s.snapshot
}

func (s *Session) Update(fn func(*Snapshot)) {

	s.mu.Lock()
	defer s.mu.Unlock()

	fn(&s.snapshot)

	s.snapshot.Updated = time.Now()

	if s.snapshot.Started.IsZero() {
		s.snapshot.Started = s.snapshot.Updated
	}
}
