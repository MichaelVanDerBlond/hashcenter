package hashcat

import (
	"strconv"
	"strings"
)

func SpeedHPS(s string) float64 {

	fields := strings.Fields(s)
	if len(fields) < 2 {
		return 0
	}

	value, err := strconv.ParseFloat(fields[0], 64)
	if err != nil {
		return 0
	}

	switch strings.ToUpper(fields[1]) {

	case "H/S":
		return value

	case "KH/S":
		return value * 1_000

	case "MH/S":
		return value * 1_000_000

	case "GH/S":
		return value * 1_000_000_000

	case "TH/S":
		return value * 1_000_000_000_000
	}

	return value
}
