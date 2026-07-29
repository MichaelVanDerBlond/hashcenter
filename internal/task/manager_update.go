package task

func (m *Manager) Update(id string, fn func(*Task)) bool {
	m.mu.Lock()
	defer m.mu.Unlock()

	t, ok := m.tasks[id]
	if !ok {
		return false
	}

	if fn != nil {
		fn(t)
	}

	return true
}

func (m *Manager) FindBySession(sessionID string) []*Task {
	m.mu.RLock()
	defer m.mu.RUnlock()

	var result []*Task

	for _, t := range m.tasks {
		if t.SessionID != sessionID {
			continue
		}

		copyTask := *t
		result = append(result, &copyTask)
	}

	return result
}
