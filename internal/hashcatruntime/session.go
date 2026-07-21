package hashcatruntime

import "sync"

type Session struct {
	mu sync.RWMutex

	snapshot Snapshot
}

func NewSession(name string) *Session {

	s := &Session{}

	s.snapshot.Session = name

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
}
