package hashcat

import "testing"

func TestPercent(t *testing.T) {

	tests := []struct {
		line string
		want float64
	}{
		{"100/100 (100.00%)", 100},
		{"5/10 (50.00%)", 50},
		{"2/20 (10.00%)", 10},
		{"0/1 (0.00%)", 0},
		{"invalid", 0},
	}

	for _, tt := range tests {

		got := percent(tt.line)

		if got != tt.want {
			t.Fatalf("%q -> %.2f, want %.2f", tt.line, got, tt.want)
		}
	}
}
