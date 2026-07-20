package hashcat

import (
	"regexp"
	"strconv"
)

var percentRE = regexp.MustCompile(`\(([0-9]+(?:\.[0-9]+)?)%\)`)

func percent(line string) float64 {

	m := percentRE.FindStringSubmatch(line)
	if len(m) != 2 {
		return 0
	}

	v, err := strconv.ParseFloat(m[1], 64)
	if err != nil {
		return 0
	}

	return v
}
