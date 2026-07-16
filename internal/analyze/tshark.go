package analyze

import (
	"strconv"
	"strings"
)

func tsharkCount(path, filter string) uint64 {

	out, err := run(
		"tshark",
		"-r", path,
		"-Y", filter,
		"-T", "fields",
		"-e", "frame.number",
	)

	if err != nil {
		return 0
	}

	lines := strings.Split(strings.TrimSpace(string(out)), "\n")

	if len(lines) == 1 && lines[0] == "" {
		return 0
	}

	return uint64(len(lines))
}

func collectTShark(r *Report) error {

	if !exists("tshark") {
		return nil
	}

	r.TShark = true

	out, err := run("tshark", "--version")
	if err == nil {

		lines := strings.Split(string(out), "\n")

		for _, line := range lines {
			line = strings.TrimSpace(line)

			if strings.HasPrefix(line, "TShark") ||
				strings.HasPrefix(line, "Wireshark") {

				r.Version = line
				break
			}
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

	if err == nil {

		lines := strings.Split(strings.TrimSpace(string(out)), "\n")

		if len(lines) == 1 && lines[0] == "" {
			r.Frames = "0"
		} else {
			r.Frames = strconv.Itoa(len(lines))
		}
	}

	r.BeaconFrames = tsharkCount(r.Path, "wlan.fc.type_subtype==8")
	r.ProbeFrames = tsharkCount(r.Path, "wlan.fc.type_subtype==4")
	r.EAPOLFrames = tsharkCount(r.Path, "eapol")

	return nil
}
