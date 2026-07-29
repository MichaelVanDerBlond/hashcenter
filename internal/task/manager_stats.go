package task

type Stats struct {
	Total    int
	Queued   int
	Running  int
	Finished int
	Failed   int
}

func (m *Manager) Stats() Stats {
	m.mu.RLock()
	defer m.mu.RUnlock()

	var s Stats

	for _, t := range m.tasks {
		s.Total++

		switch t.State {
		case Queued:
			s.Queued++
		case Running:
			s.Running++
		case Finished:
			s.Finished++
		case Failed:
			s.Failed++
		}
	}

	return s
}
