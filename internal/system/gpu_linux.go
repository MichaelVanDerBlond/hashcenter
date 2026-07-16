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

	lines := strings.Split(strings.TrimSpace(string(out)), "\n")
	if len(lines) == 0 || strings.TrimSpace(lines[0]) == "" {
		return nil
	}

	fields := strings.Split(lines[0], ",")
	if len(fields) < 6 {
		return nil
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
