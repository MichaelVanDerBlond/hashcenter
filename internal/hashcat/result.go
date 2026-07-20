package hashcat

import "time"

type Result struct {
	// Runtime

	Started time.Time `json:"started"`

	Finished time.Time `json:"finished"`

	Duration time.Duration `json:"duration"`

	ExitCode int `json:"exit_code"`

	Stdout string `json:"stdout"`

	Stderr string `json:"stderr"`

	// Parsed statistics

	State string `json:"state"`

	Speed string `json:"speed"`

	SpeedHPS float64 `json:"speed_hps"`

	Progress string `json:"progress"`

	Recovered string `json:"recovered"`

	ETA string `json:"eta"`
}
