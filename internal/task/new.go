package task

import (
	"time"

	"github.com/MichaelVanDerBlond/hashcenter/internal/attack"
	"github.com/google/uuid"
)

func NewAttack(sessionID string, job attack.Job) *Task {

	return &Task{
		ID:        uuid.NewString(),
		SessionID: sessionID,
		Type:      Attack,
		State:     Queued,
		Created:   time.Now(),

		HashFile:   job.HashFile,
		Dictionary: job.Dictionary,

		HashMode:   job.HashMode,
		AttackMode: job.AttackMode,

		Rule:        job.Rule,
		Mask:        job.Mask,
		Device:      job.Device,
		Workload:    job.Workload,
		SessionName: job.SessionName,
	}
}

func New(sessionID string, typ Type) *Task {

	return &Task{
		ID:        uuid.NewString(),
		SessionID: sessionID,
		Type:      typ,
		State:     Queued,
		Created:   time.Now(),
	}
}
