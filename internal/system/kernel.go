package system

import (
	"os/exec"
	"strings"
)

func collectKernel(info *Info) error {
	out, err := exec.Command("uname", "-r").Output()
	if err != nil {
		return err
	}

	info.Kernel = strings.TrimSpace(string(out))
	return nil
}
