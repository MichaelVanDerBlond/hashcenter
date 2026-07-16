package analyze

import (
	"os/exec"
	"strings"
)

func collectTShark(r *Report) error {

	if _, err := exec.LookPath("tshark"); err != nil {
		return nil
	}

	r.TShark = true

	out, err := exec.Command("tshark", "--version").Output()
	if err != nil {
		return nil
	}

	lines := strings.Split(string(out), "\n")

	if len(lines) > 0 {
		r.Version = strings.TrimSpace(lines[0])
	}

	return nil
}
