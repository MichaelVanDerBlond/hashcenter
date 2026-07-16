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

	text := strings.TrimSpace(string(out))
	if text == "" {
		return 0
	}

	return uint64(len(strings.Split(text, "\n")))
}

func collectTShark(r *Report) error {

	if !exists("tshark") {
		return nil
	}

	r.TShark = true

	out, _ := run("tshark", "--version")

	for _, line := range strings.Split(string(out), "\n") {

		line = strings.TrimSpace(line)

		if strings.HasPrefix(line, "TShark") ||
			strings.HasPrefix(line, "Wireshark") {

			r.Version = line
			break
		}
	}

	out, err := run(
		"tshark",
		"-r",
		r.Path,
		"-T",
		"fields",
		"-e",
		"frame.number",
	)

	if err == nil {

		text := strings.TrimSpace(string(out))

		if text == "" {
			r.Frames = "0"
		} else {
			r.Frames = strconv.Itoa(len(strings.Split(text, "\n")))
		}
	}

	// Wi-Fi counters
	r.BeaconFrames = tsharkCount(r.Path, "wlan.fc.type_subtype == 8")

	return nil
}
