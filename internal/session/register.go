package session

func (m *Manager) Register(r *Runtime) {

	if r == nil {
		return
	}

	if r.Session == nil {
		return
	}

	m.mu.Lock()
	defer m.mu.Unlock()

	m.sessions[r.Session.ID] = r
}
