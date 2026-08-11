package task

import (
	"database/sql"
	"sync"

	"github.com/MichaelVanDerBlond/hashcenter/internal/database"
)

type Manager struct {
	mu sync.RWMutex

	tasks map[string]*Task
	store Store
}

func NewManager() *Manager {
	return NewManagerWithStore(NewMemoryStore())
}

func NewManagerWithStore(store Store) *Manager {
	if store == nil {
		store = NewMemoryStore()
	}

	manager := &Manager{
		tasks: make(map[string]*Task),
		store: store,
	}

	tasks, err := store.Load()
	if err == nil {
		for _, t := range tasks {
			if t == nil || t.ID == "" {
				continue
			}

			copyTask := *t

			// A process restart means an in-memory worker no longer exists.
			// Tasks that were running must become queued again.
			if copyTask.State == Running {
				copyTask.State = Queued
			}

			manager.tasks[copyTask.ID] = &copyTask

			if copyTask.State == Queued && t.State == Running {
				_ = store.Update(&copyTask)
			}
		}
	}

	return manager
}

func NewManagerWithDB(db *sql.DB) *Manager {
	if db == nil {
		return NewManager()
	}

	return NewManagerWithStore(NewSQLiteStore(db))
}

var (
	defaultManager *Manager
	once           sync.Once
)

func DefaultManager() *Manager {
	once.Do(func() {
		db, err := database.Open()
		if err != nil {
			defaultManager = NewManager()
			return
		}

		defaultManager = NewManagerWithDB(db)
	})

	return defaultManager
}
