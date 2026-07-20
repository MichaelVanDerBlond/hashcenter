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

func (s *Status) SetState(v string) {

	s.mu.Lock()
	defer s.mu.Unlock()

	s.State = value(v)
	s.LastUpdate = time.Now()
}

func (s *Status) SetSpeed(v string) {

	s.mu.Lock()
	defer s.mu.Unlock()

	s.Speed = value(v)
	s.LastUpdate = time.Now()
}

func (s *Status) SetProgress(v string) {

	s.mu.Lock()
	defer s.mu.Unlock()

	s.Progress = value(v)
	s.LastUpdate = time.Now()
}

func (s *Status) SetRecovered(v string) {

	s.mu.Lock()
	defer s.mu.Unlock()

	s.Recovered = value(v)
	s.LastUpdate = time.Now()
}

func (s *Status) SetETA(v string) {

	s.mu.Lock()
	defer s.mu.Unlock()

	s.ETA = value(v)
	s.LastUpdate = time.Now()
}
