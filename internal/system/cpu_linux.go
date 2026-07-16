package system

import (
	"os"
	"strconv"
	"strings"
)

func collectCPU(info *Info) error {

	data, err := os.ReadFile("/proc/cpuinfo")
	if err != nil {
		return err
	}

	for _, line := range strings.Split(string(data), "\n") {

		if strings.HasPrefix(line, "model name") && info.CPU.Model == "" {
			info.CPU.Model = strings.TrimSpace(strings.Split(line, ":")[1])
		}

		if strings.HasPrefix(line, "cpu cores") && info.CPU.Cores == 0 {
			v := strings.TrimSpace(strings.Split(line, ":")[1])
			info.CPU.Cores, _ = strconv.Atoi(v)
		}
	}

	return nil
}
