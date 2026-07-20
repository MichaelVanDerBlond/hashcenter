package jobs

import "testing"

func TestManagerCreate(t *testing.T) {
	if New() == nil {
		t.Fatal("manager is nil")
	}
}
