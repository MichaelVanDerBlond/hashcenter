package session

import "time"

type Session struct {
	ID string

	State State

	Created  time.Time
	Started  time.Time
	Finished time.Time

	Interface string
	Channel   int

	Backend string

	CaptureFile string

	Error string
}

func (s *Session) IsRunning() bool {
	return s != nil && s.State == Running
}

func (s *Session) IsFinished() bool {
	return s != nil && s.State == Finished
}

func (s *Session) IsFailed() bool {
	return s != nil && s.State == Failed
}

func (s *Session) Duration() time.Duration {
	if s == nil || s.Started.IsZero() {
		return 0
	}

	end := s.Finished
	if end.IsZero() {
		end = time.Now()
	}

	return end.Sub(s.Started)
}
