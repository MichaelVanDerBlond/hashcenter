package session

func (m *Manager) List() []*Runtime {
	if m == nil {
		return nil
	}

	m.mu.RLock()
	defer m.mu.RUnlock()

	list := make([]*Runtime, 0, len(m.sessions))

	for _, runtime := range m.sessions {
		list = append(list, runtime)
	}

	return list
}
