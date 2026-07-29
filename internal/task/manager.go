package task

import "sync"

type Manager struct {
	mu sync.RWMutex

	tasks map[string]*Task
}

var (
	defaultManager *Manager
	once           sync.Once
)

func DefaultManager() *Manager {
	once.Do(func() {
		defaultManager = &Manager{
			tasks: make(map[string]*Task),
		}
	})

	return defaultManager
}

func (m *Manager) Add(t *Task) {
	if t == nil {
		return
	}

	m.mu.Lock()
	defer m.mu.Unlock()

	m.tasks[t.ID] = t
}

func (m *Manager) Get(id string) (*Task, bool) {
	m.mu.RLock()
	defer m.mu.RUnlock()

	t, ok := m.tasks[id]
	if !ok {
		return nil, false
	}

	copyTask := *t
	return &copyTask, true
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
