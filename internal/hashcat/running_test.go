package hashcat

import "testing"

func TestRunnerInitiallyStopped(t *testing.T) {

	r := New()

	if r.Running() {
		t.Fatal("runner should not be running")
	}

	if r.PID() != 0 {
		t.Fatal("pid should be zero")
	}
}
