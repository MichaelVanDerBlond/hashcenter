package task

func (m *Manager) filterByState(state State) []*Task {
	m.mu.RLock()
	defer m.mu.RUnlock()

	result := make([]*Task, 0)

	for _, t := range m.tasks {
		if t.State != state {
			continue
		}

		copyTask := *t
		result = append(result, &copyTask)
	}

	return result
}

func (m *Manager) Queued() []*Task {
	return m.filterByState(Queued)
}

func (m *Manager) Running() []*Task {
	return m.filterByState(Running)
}

func (m *Manager) Finished() []*Task {
	return m.filterByState(Finished)
}
