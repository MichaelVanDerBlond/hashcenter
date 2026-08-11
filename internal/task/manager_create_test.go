package task

import "testing"

func TestManagerTaskLifecycle(t *testing.T) {
	m := NewManager()

	task := New("session-test", Attack)

	created, err := m.Create(task)
	if err != nil {
		t.Fatal(err)
	}

	if created == nil {
		t.Fatal("expected created task")
	}

	if created.State != Queued {
		t.Fatalf("expected queued, got %s", created.State)
	}

	next, ok := m.NextQueued()
	if !ok {
		t.Fatal("expected queued task")
	}

	if next.ID != created.ID {
		t.Fatalf("task ID mismatch: %s != %s", next.ID, created.ID)
	}

	if next.State != Running {
		t.Fatalf("expected running, got %s", next.State)
	}

	if !m.SetState(next.ID, Finished) {
		t.Fatal("failed to finish task")
	}

	snapshot := m.Snapshot()

	result, ok := snapshot[next.ID]
	if !ok {
		t.Fatal("task missing from snapshot")
	}

	if result.State != Finished {
		t.Fatalf("expected finished, got %s", result.State)
	}

	if result.Started.IsZero() {
		t.Fatal("started timestamp is empty")
	}

	if result.Finished.IsZero() {
		t.Fatal("finished timestamp is empty")
	}
}
