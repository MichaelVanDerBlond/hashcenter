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
