package task

func (m *Manager) Queued() []*Task {
	m.mu.RLock()
	defer m.mu.RUnlock()

	result := make([]*Task, 0)

	for _, t := range m.tasks {
		if t.State != Queued {
			continue
		}

		copyTask := *t
		result = append(result, &copyTask)
	}

	return result
}

func (m *Manager) Running() []*Task {
	m.mu.RLock()
	defer m.mu.RUnlock()

	result := make([]*Task, 0)

	for _, t := range m.tasks {
		if t.State != Running {
			continue
		}

		copyTask := *t
		result = append(result, &copyTask)
	}

	return result
}

func (m *Manager) Finished() []*Task {
	m.mu.RLock()
	defer m.mu.RUnlock()

	result := make([]*Task, 0)

	for _, t := range m.tasks {
		if t.State != Finished {
			continue
		}

		copyTask := *t
		result = append(result, &copyTask)
	}

	return result
}
