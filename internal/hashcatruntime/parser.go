package hashcatruntime

import (
	"encoding/json"
	"time"
)

type statusJSON struct {
	Session   string  `json:"session"`
	Status    string  `json:"status"`
	Progress  float64 `json:"progress"`
	Speed     float64 `json:"speed"`
	Recovered int64   `json:"recovered"`
	Rejected  int64   `json:"rejected"`
	ETA       int64   `json:"eta"`
}

func ParseStatus(data []byte) (Snapshot, error) {

	var src statusJSON

	if err := json.Unmarshal(data, &src); err != nil {
		return Snapshot{}, err
	}

	now := time.Now()

	return Snapshot{
		Session:   src.Session,
		Status:    src.Status,
		Progress:  src.Progress,
		Speed:     src.Speed,
		Recovered: src.Recovered,
		Rejected:  src.Rejected,
		ETA:       time.Duration(src.ETA) * time.Second,
		Updated:   now,
	}, nil
}
