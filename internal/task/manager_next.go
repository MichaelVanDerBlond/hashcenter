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

		if m.store != nil {
			if err := m.store.Update(t); err != nil {
				t.State = Queued
				t.Started = t.Started.Add(0)
				return nil, false
			}
		}

		copyTask := *t
		return &copyTask, true
	}

	return nil, false
}
