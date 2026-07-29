package task

func (m *Manager) RemoveSession(sessionID string) int {
	m.mu.Lock()
	defer m.mu.Unlock()

	removed := 0

	for id, t := range m.tasks {
		if t.SessionID != sessionID {
			continue
		}

		delete(m.tasks, id)
		removed++
	}

	return removed
}

func (m *Manager) HasRunningSession(sessionID string) bool {
	m.mu.RLock()
	defer m.mu.RUnlock()

	for _, t := range m.tasks {
		if t.SessionID == sessionID && t.State == Running {
			return true
		}
	}

	return false
}
