package analyze

import (
	"strings"
)

func collectTShark(r *Report) error {

	if !exists("tshark") {
		return nil
	}

	r.TShark = true

	out, err := run("tshark", "--version")
	if err == nil {

		lines := strings.Split(string(out), "\n")

		if len(lines) > 0 {
			r.Version = strings.TrimSpace(lines[0])
		}
	}

	out, err = run(
		"tshark",
		"-r",
		r.Path,
		"-T",
		"fields",
		"-e",
		"frame.number",
	)

	if err != nil {
		return nil
	}

	lines := strings.Split(strings.TrimSpace(string(out)), "\n")

	if len(lines) == 1 && lines[0] == "" {
		r.Frames = "0"
	} else {
		r.Frames = strings.TrimSpace(lines[len(lines)-1])
	}

	return nil
}
