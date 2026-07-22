package hashcat

import (
	"strings"
	"time"
)

func value(line string) string {
	if i := strings.Index(line, ":"); i >= 0 {
		return strings.TrimSpace(line[i+1:])
	}

	return strings.TrimSpace(line)
}

func (s *Status) set(update *string, v string) {
	s.mu.Lock()
	defer s.mu.Unlock()

	*update = value(v)
	s.LastUpdate = time.Now()
}

func (s *Status) SetState(v string) {
	s.set(&s.State, v)
}

func (s *Status) SetSpeed(v string) {
	s.set(&s.Speed, v)
}

func (s *Status) SetProgress(v string) {
	s.set(&s.Progress, v)
}

func (s *Status) SetRecovered(v string) {
	s.set(&s.Recovered, v)
}

func (s *Status) SetETA(v string) {
	s.set(&s.ETA, v)
}
