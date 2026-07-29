package task

import "sync"

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

	return &Manager{
		tasks: make(map[string]*Task),
		store: store,
	}
}

var (
	defaultManager *Manager
	once           sync.Once
)

func DefaultManager() *Manager {
	once.Do(func() {
		defaultManager = NewManager()
	})

	return defaultManager
}
