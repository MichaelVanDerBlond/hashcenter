package status

import "testing"

func TestStatusStruct(t *testing.T) {

	s := &Status{}

	if s == nil {
		t.Fatal("status is nil")
	}
}
