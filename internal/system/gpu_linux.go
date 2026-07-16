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

	fields := strings.Split(strings.TrimSpace(string(out)), ",")

	if len(fields) >= 6 {
		info.GPU.Name = strings.TrimSpace(fields[0])
		info.GPU.Driver = strings.TrimSpace(fields[1])
		info.GPU.MemoryTotal = strings.TrimSpace(fields[2])
		info.GPU.MemoryUsed = strings.TrimSpace(fields[3])
		info.GPU.Temperature = strings.TrimSpace(fields[4])
		info.GPU.Utilization = strings.TrimSpace(fields[5])
	}

	return nil
}
