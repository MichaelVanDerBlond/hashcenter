package hashcat

func (s *Status) Stats() Stats {

	s.mu.RLock()
	defer s.mu.RUnlock()

	return Stats{
		State:            s.State,
		Speed:            s.Speed,
		SpeedHPS:         SpeedHPS(s.Speed),
		Progress:         s.Progress,
		Recovered:        s.Recovered,
		ETA:              s.ETA,
		ProgressPercent:  percent(s.Progress),
		RecoveredPercent: percent(s.Recovered),
	}
}
