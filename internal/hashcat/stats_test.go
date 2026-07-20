package hashcat

import "testing"

func TestStatusStats(t *testing.T) {

	s := NewStatus()

	s.SetState("Status...........: Running")
	s.SetSpeed("Speed.#1.........: 150 kH/s")
	s.SetProgress("Progress.........: 50/100 (50.00%)")
	s.SetRecovered("Recovered........: 2/10 (20.00%)")
	s.SetETA("Time.Estimated...: tomorrow")

	stats := s.Stats()

	if stats.State != "Running" {
		t.Fatal(stats.State)
	}

	if stats.Speed != "150 kH/s" {
		t.Fatal(stats.Speed)
	}

	if stats.Progress != "50/100 (50.00%)" {
		t.Fatal(stats.Progress)
	}

	if stats.ProgressPercent != 50 {
		t.Fatal(stats.ProgressPercent)
	}

	if stats.Recovered != "2/10 (20.00%)" {
		t.Fatal(stats.Recovered)
	}

	if stats.RecoveredPercent != 20 {
		t.Fatal(stats.RecoveredPercent)
	}

	if stats.ETA != "tomorrow" {
		t.Fatal(stats.ETA)
	}
}
