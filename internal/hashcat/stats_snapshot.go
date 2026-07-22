package hashcat

func (s *Status) Stats() Stats {
	snapshot := s.Snapshot()

	return Stats{
		State:            snapshot.State,
		Speed:            snapshot.Speed,
		SpeedHPS:         SpeedHPS(snapshot.Speed),
		Progress:         snapshot.Progress,
		Recovered:        snapshot.Recovered,
		ETA:              snapshot.ETA,
		ProgressPercent:  percent(snapshot.Progress),
		RecoveredPercent: percent(snapshot.Recovered),
	}
}
