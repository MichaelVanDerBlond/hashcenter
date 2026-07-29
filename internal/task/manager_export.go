package task

import (
	"sort"
)

func (m *Manager) All() []*Task {
	m.mu.RLock()
	defer m.mu.RUnlock()

	result := make([]*Task, 0, len(m.tasks))

	for _, t := range m.tasks {
		if t == nil {
			continue
		}

		cp := *t
		result = append(result, &cp)
	}

	sort.Slice(result, func(i, j int) bool {
		return result[i].Created.Before(result[j].Created)
	})

	return result
}
