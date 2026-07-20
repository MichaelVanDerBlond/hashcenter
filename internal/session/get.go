package session

func (m *Manager) Get(id string) (*Runtime, bool) {

	m.mu.RLock()
	defer m.mu.RUnlock()

	r, ok := m.sessions[id]

	return r, ok
}
