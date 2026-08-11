package services

import (
	"github.com/MichaelVanDerBlond/hashcenter/internal/hashcat"
	"github.com/MichaelVanDerBlond/hashcenter/internal/jobs"
)

type RuntimeSnapshot struct {
	SessionID string `json:"session_id"`

	State string `json:"state"`

	Speed string `json:"speed"`

	SpeedHPS float64 `json:"speed_hps"`

	Progress string `json:"progress"`

	ProgressPercent float64 `json:"progress_percent"`

	Recovered string `json:"recovered"`

	RecoveredPercent float64 `json:"recovered_percent"`

	ETA string `json:"eta"`

	LastUpdate string `json:"last_update"`
}

func ListRuntime() []RuntimeSnapshot {
	runtimes := jobs.DefaultManager().List()

	result := make([]RuntimeSnapshot, 0, len(runtimes))

	for _, runtime := range runtimes {
		if runtime == nil {
			continue
		}

		snapshot := RuntimeSnapshot{
			SessionID: runtime.SessionID,
		}

		if runtime.Status != nil {
			stats := runtime.Status.Stats()

			snapshot.State = stats.State
			snapshot.Speed = stats.Speed
			snapshot.SpeedHPS = stats.SpeedHPS
			snapshot.Progress = stats.Progress
			snapshot.ProgressPercent = stats.ProgressPercent
			snapshot.Recovered = stats.Recovered
			snapshot.RecoveredPercent = stats.RecoveredPercent
			snapshot.ETA = stats.ETA

			status := runtime.Status.Snapshot()
			snapshot.LastUpdate = status.LastUpdate.Format("2006-01-02 15:04:05")
		}

		result = append(result, snapshot)
	}

	return result
}

var _ = hashcat.Stats{}
