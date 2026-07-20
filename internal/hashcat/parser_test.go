package hashcat

import "testing"

func TestParserRoutesFields(t *testing.T) {

	p := NewParser()

	p.Parse("Status...........: Running")
	p.Parse("Speed.#1.........: 12345 H/s")
	p.Parse("Progress.........: 100/100 (100.00%)")
	p.Parse("Recovered........: 1/1 (100.00%)")
	p.Parse("Time.Estimated...: Mon Jan 01 00:00:00 2026 (0 secs)")

	s := p.Status.Snapshot()

	if s.State != "Running" {
		t.Fatalf("unexpected state: %q", s.State)
	}

	if s.Speed != "12345 H/s" {
		t.Fatalf("unexpected speed: %q", s.Speed)
	}

	if s.Progress != "100/100 (100.00%)" {
		t.Fatalf("unexpected progress: %q", s.Progress)
	}

	if s.Recovered != "1/1 (100.00%)" {
		t.Fatalf("unexpected recovered: %q", s.Recovered)
	}

	if s.ETA != "Mon Jan 01 00:00:00 2026 (0 secs)" {
		t.Fatalf("unexpected eta: %q", s.ETA)
	}
}

func TestParserIgnoresUnknownLines(t *testing.T) {

	p := NewParser()

	before := p.Status.Snapshot()

	p.Parse("")
	p.Parse("   ")
	p.Parse("Some random text")
	p.Parse("Another line")

	after := p.Status.Snapshot()

	if before != after {
		t.Fatal("status changed unexpectedly")
	}
}
