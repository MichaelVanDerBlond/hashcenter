package web

import (
	"time"

	"github.com/MichaelVanDerBlond/hashcenter/internal/task"
)

type QueueJob struct {
	ID         string    `json:"id"`
	HashFile   string    `json:"hash_file"`
	Dictionary string    `json:"dictionary"`
	State      string    `json:"state"`
	Started    time.Time `json:"started,omitempty"`
	Finished   time.Time `json:"finished,omitempty"`
	Error      string    `json:"error,omitempty"`
}

func listQueueJobs() []*QueueJob {
	tasks := task.DefaultManager().All()

	result := make([]*QueueJob, 0, len(tasks))

	for _, t := range tasks {
		if t == nil || t.Type != task.Attack {
			continue
		}

		result = append(result, &QueueJob{
			ID:         t.ID,
			HashFile:   t.HashFile,
			Dictionary: t.Dictionary,
			State:      string(t.State),
			Started:    t.Started,
			Finished:   t.Finished,
			Error:      t.Error,
		})
	}

	return result
}
