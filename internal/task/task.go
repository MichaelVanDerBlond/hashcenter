package task

import "time"

type Task struct {
	ID string

	SessionID string

	Type Type

	State State

	Created time.Time

	Started time.Time

	Finished time.Time

	Error string
}
