package hashcat

import "testing"

func TestSpeedHPS(t *testing.T) {

	tests := []struct {
		in   string
		want float64
	}{
		{"100 H/s", 100},
		{"1.5 kH/s", 1500},
		{"2 MH/s", 2000000},
		{"3 GH/s", 3000000000},
		{"0.5 TH/s", 500000000000},
		{"invalid", 0},
	}

	for _, tt := range tests {
		got := SpeedHPS(tt.in)

		if got != tt.want {
			t.Fatalf("%q -> %.0f, want %.0f", tt.in, got, tt.want)
		}
	}
}
