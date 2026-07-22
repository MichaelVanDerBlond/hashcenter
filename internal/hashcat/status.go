package hashcat

import (
	"sync"
	"time"
)

type Status struct {
	mu sync.RWMutex

	State string

	Speed string

	Progress string

	Recovered string

	ETA string

	LastUpdate time.Time
}

func NewStatus() *Status {
	return &Status{
		State: "starting",
	}
}

func (s *Status) Snapshot() Status {
	s.mu.RLock()
	defer s.mu.RUnlock()

	return s.snapshotLocked()
}

func (s *Status) snapshotLocked() Status {
	return Status{
		State:      s.State,
		Speed:      s.Speed,
		Progress:   s.Progress,
		Recovered:  s.Recovered,
		ETA:        s.ETA,
		LastUpdate: s.LastUpdate,
	}
}
