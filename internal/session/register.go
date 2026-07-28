package session

func (m *Manager) Register(r *Runtime) {
	if m == nil || r == nil || r.Session == nil {
		return
	}

	if r.Session.ID == "" {
		return
	}

	m.mu.Lock()
	defer m.mu.Unlock()

	m.sessions[r.Session.ID] = r
}
