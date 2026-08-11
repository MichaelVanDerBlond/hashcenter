package task

func (m *Manager) Create(t *Task) (*Task, error) {
	if m == nil || t == nil {
		return nil, nil
	}

	m.mu.Lock()
	defer m.mu.Unlock()

	if t.ID == "" {
		return nil, nil
	}

	if _, exists := m.tasks[t.ID]; exists {
		return nil, nil
	}

	copyTask := *t
	m.tasks[t.ID] = &copyTask

	if m.store != nil {
		if err := m.store.Save(&copyTask); err != nil {
			delete(m.tasks, t.ID)
			return nil, err
		}
	}

	result := copyTask

	return &result, nil
}
