package hashcat

type Stats struct {
	State string `json:"state"`

	Speed string `json:"speed"`

	SpeedHPS float64 `json:"speed_hps"`

	Progress string `json:"progress"`

	Recovered string `json:"recovered"`

	ETA string `json:"eta"`

	ProgressPercent float64 `json:"progress_percent"`

	RecoveredPercent float64 `json:"recovered_percent"`
}
