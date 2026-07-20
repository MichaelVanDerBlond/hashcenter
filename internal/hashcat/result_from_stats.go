package hashcat

import "time"

func NewResult(stats Stats, started time.Time) *Result {

	return &Result{
		State:     stats.State,
		Speed:     stats.Speed,
		SpeedHPS:  stats.SpeedHPS,
		Progress:  stats.Progress,
		Recovered: stats.Recovered,
		ETA:       stats.ETA,
		Duration:  time.Since(started),
	}
}
