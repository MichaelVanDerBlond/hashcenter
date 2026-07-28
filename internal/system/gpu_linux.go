package system

import (
	"os/exec"
	"strings"
)

func collectGPU(info *Info) error {

	out, err := exec.Command(
		"nvidia-smi",
		"--query-gpu=name,driver_version,memory.total,memory.used,temperature.gpu,utilization.gpu",
		"--format=csv,noheader,nounits",
	).Output()
	if err != nil {
		return nil
	}

	for _, line := range strings.Split(strings.TrimSpace(string(out)), "\n") {

		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}

		fields := strings.Split(line, ",")
		if len(fields) < 6 {
			continue
		}

		info.GPU = GPUInfo{
			Name:        strings.TrimSpace(fields[0]),
			Driver:      strings.TrimSpace(fields[1]),
			MemoryTotal: strings.TrimSpace(fields[2]),
			MemoryUsed:  strings.TrimSpace(fields[3]),
			Temperature: strings.TrimSpace(fields[4]),
			Utilization: strings.TrimSpace(fields[5]),
		}

		return nil
	}

	return nil
}
