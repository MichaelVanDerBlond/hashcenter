package task

func (m *Manager) Remove(id string) bool {
	m.mu.Lock()
	defer m.mu.Unlock()

	if _, ok := m.tasks[id]; !ok {
		return false
	}

	delete(m.tasks, id)

	return true
}

func (m *Manager) Count() int {
	m.mu.RLock()
	defer m.mu.RUnlock()

	return len(m.tasks)
}

func (m *Manager) Exists(id string) bool {
	m.mu.RLock()
	defer m.mu.RUnlock()

	_, ok := m.tasks[id]
	return ok
}
