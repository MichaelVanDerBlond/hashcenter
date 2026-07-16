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
	if err == nil {

		lines := strings.Split(string(out), "\n")

		if len(lines) > 0 {
			r.Version = strings.TrimSpace(lines[0])
		}
	}

	out, err = exec.Command(
		"tshark",
		"-r",
		r.Path,
		"-T",
		"fields",
		"-e",
		"frame.number",
	).Output()

	if err != nil {
		return nil
	}

	lines := strings.Split(strings.TrimSpace(string(out)), "\n")

	if len(lines) == 1 && lines[0] == "" {
		r.Frames = "0"
	} else {
		r.Frames = strings.TrimSpace(
			strings.Split(strings.TrimSpace(string(out)), "\n")[len(lines)-1],
		)
	}

	return nil
}
