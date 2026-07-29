package task

import (
	"time"

	"github.com/google/uuid"
)

func New(sessionID string, typ Type) *Task {

	return &Task{
		ID:         uuid.NewString(),
		SessionID:  sessionID,
		Type:       typ,
		State:      Queued,
		Created:    time.Now(),
		HashMode:   22000,
		AttackMode: 0,
		Workload:   3,
	}
}
