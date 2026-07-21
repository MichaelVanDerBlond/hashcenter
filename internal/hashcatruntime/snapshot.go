package hashcatruntime

import "time"

type Snapshot struct {
	Session string `json:"session"`

	Status string `json:"status"`

	Progress float64 `json:"progress"`

	Speed float64 `json:"speed"`

	Recovered int64 `json:"recovered"`

	Rejected int64 `json:"rejected"`

	ETA time.Duration `json:"eta"`

	Started time.Time `json:"started"`

	Updated time.Time `json:"updated"`
}
