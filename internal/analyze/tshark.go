package analyze

import (
	"strconv"
	"strings"
)

func tsharkCount(path, filter string) uint64 {

	args := []string{
		"-r", path,
	}

	if filter != "" {
		args = append(args,
			"-Y", filter,
		)
	}

	args = append(args,
		"-T", "fields",
		"-e", "frame.number",
	)

	out, err := run("tshark", args...)
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

	r.Frames = strconv.FormatUint(tsharkCount(r.Path, ""), 10)

	r.BeaconFrames = tsharkCount(r.Path, "wlan.fc.type_subtype == 8")
	r.ProbeFrames = tsharkCount(r.Path, "wlan.fc.type_subtype == 4")

	return nil
}
