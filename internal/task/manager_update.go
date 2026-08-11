package task

func (m *Manager) Update(id string, fn func(*Task)) bool {
	if m == nil || id == "" {
		return false
	}

	m.mu.Lock()
	defer m.mu.Unlock()

	t, ok := m.tasks[id]
	if !ok || t == nil {
		return false
	}

	if fn != nil {
		fn(t)
	}

	if m.store != nil {
		if err := m.store.Update(t); err != nil {
			return false
		}
	}

	return true
}
