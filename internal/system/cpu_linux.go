package system

import (
	"os"
	"runtime"
	"strconv"
	"strings"
)

func collectCPU(info *Info) error {

	info.CPU.Arch = runtime.GOARCH

	data, err := os.ReadFile("/proc/cpuinfo")
	if err != nil {
		return err
	}

	for _, line := range strings.Split(string(data), "\n") {

		parts := strings.SplitN(line, ":", 2)
		if len(parts) != 2 {
			continue
		}

		key := strings.TrimSpace(parts[0])
		value := strings.TrimSpace(parts[1])

		switch {

		case key == "vendor_id" && info.CPU.Vendor == "":
			info.CPU.Vendor = value

		case key == "model name" && info.CPU.Model == "":
			info.CPU.Model = value

		case key == "cpu cores" && info.CPU.Cores == 0:
			info.CPU.Cores, _ = strconv.Atoi(value)

		case key == "processor":
			info.CPU.Threads++
		}
	}

	return nil
}
