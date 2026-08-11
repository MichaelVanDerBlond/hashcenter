package task

import "time"

func (m *Manager) SetState(id string, state State) bool {
	m.mu.Lock()
	defer m.mu.Unlock()

	t, ok := m.tasks[id]
	if !ok {
		return false
	}

	t.State = state

	switch state {
	case Running:
		if t.Started.IsZero() {
			t.Started = time.Now()
		}
	case Finished, Failed:
		if t.Finished.IsZero() {
			t.Finished = time.Now()
		}
	}

	if m.store != nil {
		if err := m.store.Update(t); err != nil {
			return false
		}
	}

	return true
}

func (m *Manager) SetError(id, errText string) bool {
	m.mu.Lock()
	defer m.mu.Unlock()

	t, ok := m.tasks[id]
	if !ok {
		return false
	}

	t.Error = errText
	t.State = Failed

	if t.Finished.IsZero() {
		t.Finished = time.Now()
	}

	if m.store != nil {
		if err := m.store.Update(t); err != nil {
			return false
		}
	}

	return true
}
