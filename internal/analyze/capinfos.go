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

	out, err := exec.Command("capinfos", r.Path).CombinedOutput()
	if err != nil {
		return nil
	}

	value := func(line, prefix string) string {
		return strings.TrimSpace(strings.TrimPrefix(line, prefix))
	}

	for _, line := range strings.Split(string(out), "\n") {

		line = strings.TrimSpace(line)

		switch {

		case strings.HasPrefix(line, "Number of packets:"):
			r.Packets = value(line, "Number of packets:")

		case strings.HasPrefix(line, "Capture duration:"):
			r.Duration = value(line, "Capture duration:")

		case strings.HasPrefix(line, "File encapsulation:"):
			r.Encapsulation = value(line, "File encapsulation:")
		}
	}

	return nil
}
