package session

import "sync"

type Manager struct {
	mu sync.RWMutex

	sessions map[string]*Runtime
}

func NewManager() *Manager {
	return &Manager{
		sessions: make(map[string]*Runtime),
	}
}

func (m *Manager) Count() int {
	if m == nil {
		return 0
	}

	m.mu.RLock()
	defer m.mu.RUnlock()

	return len(m.sessions)
}
