package jobs

func (m *Manager) Register(runtime *Runtime) {
	if runtime == nil {
		return
	}

	m.mu.Lock()
	defer m.mu.Unlock()

	m.jobs[runtime.SessionID] = runtime
}

func (m *Manager) Remove(sessionID string) {
	m.mu.Lock()
	defer m.mu.Unlock()

	delete(m.jobs, sessionID)
}

func (m *Manager) Get(sessionID string) (*Runtime, bool) {
	m.mu.RLock()
	defer m.mu.RUnlock()

	runtime, ok := m.jobs[sessionID]

	return runtime, ok
}

func (m *Manager) List() []*Runtime {
	m.mu.RLock()
	defer m.mu.RUnlock()

	result := make([]*Runtime, 0, len(m.jobs))

	for _, runtime := range m.jobs {
		result = append(result, runtime)
	}

	return result
}
