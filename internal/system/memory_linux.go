package system

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func collectMemory(info *Info) error {

	f, err := os.Open("/proc/meminfo")
	if err != nil {
		return err
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {

		fields := strings.Fields(scanner.Text())
		if len(fields) < 2 {
			continue
		}

		value, err := strconv.ParseUint(fields[1], 10, 64)
		if err != nil {
			continue
		}

		value *= 1024

		switch strings.TrimSuffix(fields[0], ":") {
		case "MemTotal":
			info.Memory.Total = value
		case "MemAvailable":
			info.Memory.Available = value
		case "MemFree":
			info.Memory.Free = value
		}
	}

	return scanner.Err()
}
