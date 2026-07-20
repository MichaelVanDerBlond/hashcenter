package session

func (m *Manager) Remove(id string) {

	if id == "" {
		return
	}

	m.mu.Lock()
	defer m.mu.Unlock()

	delete(m.sessions, id)
}
