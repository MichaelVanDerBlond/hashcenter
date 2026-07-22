package hashcat

import (
	"strconv"
	"strings"
)

var speedUnits = map[string]float64{
	"H/S":  1,
	"KH/S": 1_000,
	"MH/S": 1_000_000,
	"GH/S": 1_000_000_000,
	"TH/S": 1_000_000_000_000,
}

func SpeedHPS(s string) float64 {
	fields := strings.Fields(s)
	if len(fields) < 2 {
		return 0
	}

	value, err := strconv.ParseFloat(fields[0], 64)
	if err != nil {
		return 0
	}

	if mul, ok := speedUnits[strings.ToUpper(fields[1])]; ok {
		return value * mul
	}

	return value
}
