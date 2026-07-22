package hashcat

import (
	"regexp"
	"strconv"
)

var percentRE = regexp.MustCompile(`\(([0-9]+(?:\.[0-9]+)?)%\)`)

func percent(line string) float64 {
	match := percentRE.FindStringSubmatch(line)
	if len(match) != 2 {
		return 0
	}

	value, err := strconv.ParseFloat(match[1], 64)
	if err != nil {
		return 0
	}

	return value
}
