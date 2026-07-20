package hashcat

import (
	"testing"
	"time"
)

func TestNewResult(t *testing.T) {

	stats := Stats{
		State:     "Running",
		Speed:     "250 MH/s",
		SpeedHPS:  250000000,
		Progress:  "5/10 (50.00%)",
		Recovered: "1/5 (20.00%)",
		ETA:       "5 mins",
	}

	r := NewResult(stats, time.Now().Add(-time.Second))

	if r.State != stats.State {
		t.Fatal("state mismatch")
	}

	if r.Speed != stats.Speed {
		t.Fatal("speed mismatch")
	}

	if r.SpeedHPS != stats.SpeedHPS {
		t.Fatal("speed_hps mismatch")
	}

	if r.Duration <= 0 {
		t.Fatal("duration not calculated")
	}
}
