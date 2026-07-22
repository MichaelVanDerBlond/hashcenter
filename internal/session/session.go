package session

import "time"

type Session struct {
	ID string

	State State

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
