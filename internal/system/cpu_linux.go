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

		switch {

		case strings.HasPrefix(line, "vendor_id") && info.CPU.Vendor == "":
			info.CPU.Vendor = strings.TrimSpace(strings.SplitN(line, ":", 2)[1])

		case strings.HasPrefix(line, "model name") && info.CPU.Model == "":
			info.CPU.Model = strings.TrimSpace(strings.SplitN(line, ":", 2)[1])

		case strings.HasPrefix(line, "cpu cores") && info.CPU.Cores == 0:
			v := strings.TrimSpace(strings.SplitN(line, ":", 2)[1])
			info.CPU.Cores, _ = strconv.Atoi(v)

		case strings.HasPrefix(line, "processor"):
			info.CPU.Threads++
		}
	}

	return nil
}
