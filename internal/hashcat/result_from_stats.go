package hashcat

import "time"

func NewResult(stats Stats, started time.Time) *Result {
	finished := time.Now()

	return &Result{
		Started:   started,
		Finished:  finished,
		Duration:  finished.Sub(started),
		State:     stats.State,
		Speed:     stats.Speed,
		SpeedHPS:  stats.SpeedHPS,
		Progress:  stats.Progress,
		Recovered: stats.Recovered,
		ETA:       stats.ETA,
	}
}
