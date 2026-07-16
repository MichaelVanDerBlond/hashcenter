package system

import (
	"os"
	"strings"
)

func collectOS(info *Info) error {
	data, err := os.ReadFile("/etc/os-release")
	if err != nil {
		return err
	}

	for _, line := range strings.Split(string(data), "\n") {
		if strings.HasPrefix(line, "PRETTY_NAME=") {
			info.OS = strings.Trim(strings.TrimPrefix(line, "PRETTY_NAME="), "\"")
			break
		}
	}

	return nil
}
