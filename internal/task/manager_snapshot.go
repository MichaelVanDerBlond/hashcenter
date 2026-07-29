package task

func (m *Manager) Snapshot() map[string]Task {
	m.mu.RLock()
	defer m.mu.RUnlock()

	snapshot := make(map[string]Task, len(m.tasks))

	for id, t := range m.tasks {
		if t == nil {
			continue
		}

		snapshot[id] = *t
	}

	return snapshot
}

func (m *Manager) Restore(snapshot map[string]Task) {
	m.mu.Lock()
	defer m.mu.Unlock()

	m.tasks = make(map[string]*Task, len(snapshot))

	for id, t := range snapshot {
		task := t
		m.tasks[id] = &task
	}
}
