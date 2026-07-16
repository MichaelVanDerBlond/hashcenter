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

	s := bufio.NewScanner(f)

	for s.Scan() {

		line := s.Text()

		if strings.HasPrefix(line, "MemTotal:") {

			fields := strings.Fields(line)

			if len(fields) >= 2 {

				v, _ := strconv.ParseUint(fields[1], 10, 64)

				info.Memory.Total = v * 1024
			}

			break
		}
	}

	return nil
}
