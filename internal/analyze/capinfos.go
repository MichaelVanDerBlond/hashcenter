package analyze

import (
	"os/exec"
	"strings"
)

func collectCapinfos(r *Report) error {

	if _, err := exec.LookPath("capinfos"); err != nil {
		return nil
	}

	r.Capinfos = true

	out, err := exec.Command("capinfos", r.Path).Output()
	if err != nil {
		return nil
	}

	for _, line := range strings.Split(string(out), "\n") {

		line = strings.TrimSpace(line)

		switch {

		case strings.HasPrefix(line, "Number of packets:"):
			r.Packets = strings.TrimSpace(strings.TrimPrefix(line, "Number of packets:"))

		case strings.HasPrefix(line, "Capture duration:"):
			r.Duration = strings.TrimSpace(strings.TrimPrefix(line, "Capture duration:"))

		case strings.HasPrefix(line, "File encapsulation:"):
			r.Encapsulation = strings.TrimSpace(strings.TrimPrefix(line, "File encapsulation:"))
		}
	}

	return nil
}
