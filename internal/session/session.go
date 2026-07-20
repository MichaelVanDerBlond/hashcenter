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
