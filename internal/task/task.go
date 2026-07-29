package task

import "time"

type Task struct {
	ID string

	SessionID string

	Type Type

	State State

	Created  time.Time
	Started  time.Time
	Finished time.Time

	HashFile string

	Dictionary string

	HashMode int

	AttackMode int

	Rule string

	Mask string

	Device string

	Workload int

	SessionName string

	Error string
}
