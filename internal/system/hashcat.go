package system

import (
	"os/exec"
	"strings"
)

func collectHashcat(info *Info) error {
	out, err := exec.Command("hashcat", "--version").Output()
	if err != nil {
		info.Hashcat = "not installed"
		return nil
	}

	info.Hashcat = strings.TrimSpace(string(out))
	return nil
}
