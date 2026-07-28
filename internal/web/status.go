package web

import (
	"sync"
	"time"
)

type QueueJob struct {
	ID         int       `json:"id"`
	HashFile   string    `json:"hash_file"`
	Dictionary string    `json:"dictionary"`
	State      string    `json:"state"`
	Started    time.Time `json:"started,omitempty"`
	Finished   time.Time `json:"finished,omitempty"`
}

var (
	queueMu sync.RWMutex

	queueJobs []*QueueJob

	nextQueueID int
)

func addQueueJob(hash, dict string) *QueueJob {
	queueMu.Lock()
	defer queueMu.Unlock()

	nextQueueID++

	job := &QueueJob{
		ID:         nextQueueID,
		HashFile:   hash,
		Dictionary: dict,
		State:      "queued",
	}

	queueJobs = append(queueJobs, job)

	return job
}

func listQueueJobs() []*QueueJob {
	queueMu.RLock()
	defer queueMu.RUnlock()

	result := make([]*QueueJob, 0, len(queueJobs))

	for _, job := range queueJobs {
		copyJob := *job
		result = append(result, &copyJob)
	}

	return result
}
