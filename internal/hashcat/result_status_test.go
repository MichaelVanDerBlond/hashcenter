package hashcat

import "testing"

func TestResultHelpers(t *testing.T) {

	r := &Result{
		ExitCode: 0,
		State:    "Cracked",
	}

	if !r.Success() {
		t.Fatal("expected success")
	}

	if !r.Cracked() {
		t.Fatal("expected cracked")
	}

	if r.Exhausted() {
		t.Fatal("unexpected exhausted")
	}

	if r.Running() {
		t.Fatal("unexpected running")
	}

	r.State = "Running"

	if !r.Running() {
		t.Fatal("expected running")
	}
}
