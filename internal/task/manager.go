package task

import (
	"context"
	"sync"
	"time"

	"github.com/MichaelVanDerBlond/hashcenter/internal/attack"
)

type State string

const (
	StateQueued   State = "queued"
	StateRunning  State = "running"
	StateFinished State = "finished"
	StateFailed   State = "failed"
)

type Task struct {
	ID int `json:"id"`

	HashFile   string `json:"hash_file"`
	Dictionary string `json:"dictionary"`

	State State `json:"state"`

	Created  time.Time `json:"created"`
	Started  time.Time `json:"started,omitempty"`
	Finished time.Time `json:"finished,omitempty"`

	Error string `json:"error,omitempty"`
}

type Manager struct {
	mu sync.RWMutex

	nextID int

	queue chan queuedTask

	tasks []*Task
}

type queuedTask struct {
	meta *Task
	job  attack.Job
}

var (
	defaultManager *Manager
	once           sync.Once
)

func Default() *Manager {
	once.Do(func() {
		defaultManager = &Manager{
			queue: make(chan queuedTask, 128),
		}

		go defaultManager.worker()
	})

	return defaultManager
}

func (m *Manager) worker() {
	for item := range m.queue {

		m.mu.Lock()
		item.meta.State = StateRunning
		item.meta.Started = time.Now()
		m.mu.Unlock()

		err := attack.Execute(context.Background(), item.job)

		m.mu.Lock()

		item.meta.Finished = time.Now()

		if err != nil {
			item.meta.State = StateFailed
			item.meta.Error = err.Error()
		} else {
			item.meta.State = StateFinished
		}

		m.mu.Unlock()
	}
}

func (m *Manager) Enqueue(job attack.Job) *Task {

	m.mu.Lock()

	m.nextID++

	t := &Task{
		ID:         m.nextID,
		HashFile:   job.HashFile,
		Dictionary: job.Dictionary,
		State:      StateQueued,
		Created:    time.Now(),
	}

	m.tasks = append(m.tasks, t)

	m.mu.Unlock()

	m.queue <- queuedTask{
		meta: t,
		job:  job,
	}

	return t
}

func (m *Manager) List() []*Task {

	m.mu.RLock()
	defer m.mu.RUnlock()

	result := make([]*Task, 0, len(m.tasks))

	for _, t := range m.tasks {
		copyTask := *t
		result = append(result, &copyTask)
	}

	return result
}
