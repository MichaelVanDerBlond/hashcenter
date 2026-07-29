package task

func (m *Manager) ClearFinished() int {
	m.mu.Lock()
	defer m.mu.Unlock()

	removed := 0

	for id, t := range m.tasks {
		if t.State != Finished {
			continue
		}

		delete(m.tasks, id)
		removed++
	}

	return removed
}

func (m *Manager) ClearFailed() int {
	m.mu.Lock()
	defer m.mu.Unlock()

	removed := 0

	for id, t := range m.tasks {
		if t.State != Failed {
			continue
		}

		delete(m.tasks, id)
		removed++
	}

	return removed
}

func (m *Manager) ClearCompleted() int {
	m.mu.Lock()
	defer m.mu.Unlock()

	removed := 0

	for id, t := range m.tasks {
		switch t.State {
		case Finished, Failed:
			delete(m.tasks, id)
			removed++
		}
	}

	return removed
}
