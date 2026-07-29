package task

func (m *Manager) NextQueued() (*Task, bool) {
	m.mu.Lock()
	defer m.mu.Unlock()

	for _, t := range m.tasks {
		if t.State != Queued {
			continue
		}

		t.State = Running

		if t.Started.IsZero() {
			t.Started = now()
		}

		copyTask := *t
		return &copyTask, true
	}

	return nil, false
}
