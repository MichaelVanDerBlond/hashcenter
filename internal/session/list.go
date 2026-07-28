package session

import "sort"

func (m *Manager) List() []*Runtime {
	if m == nil {
		return nil
	}

	m.mu.RLock()
	defer m.mu.RUnlock()

	list := make([]*Runtime, 0, len(m.sessions))

	for _, runtime := range m.sessions {
		list = append(list, runtime)
	}

	sort.Slice(list, func(i, j int) bool {
		if list[i] == nil || list[i].Session == nil {
			return false
		}
		if list[j] == nil || list[j].Session == nil {
			return true
		}
		return list[i].Session.ID < list[j].Session.ID
	})

	return list
}
