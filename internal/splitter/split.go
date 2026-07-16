package splitter

import (
	"encoding/hex"
	"os"
	"path/filepath"
	"strings"
)

func Split(file string) error {

	data, err := os.ReadFile(file)
	if err != nil {
		return err
	}

	if err := os.MkdirAll("output/hashes", 0755); err != nil {
		return err
	}

	networks := make(map[string][]string)

	lines := strings.Split(strings.TrimSpace(string(data)), "\n")

	for _, line := range lines {

		if strings.TrimSpace(line) == "" {
			continue
		}

		fields := strings.Split(line, "*")

		if len(fields) < 6 {
			continue
		}

		bssid := strings.ToUpper(fields[3])

		essid := fields[5]

		name := essid

		if b, err := hex.DecodeString(essid); err == nil {
			name = string(b)
		}

		name = strings.TrimSpace(name)

		if name == "" {
			name = "hidden"
		}

		name = strings.ReplaceAll(name, "/", "_")
		name = strings.ReplaceAll(name, "\\", "_")
		name = strings.ReplaceAll(name, ":", "-")

		key := bssid + "__" + name

		networks[key] = append(networks[key], line)
	}

	for key, hashes := range networks {

		filename := filepath.Join(
			"output",
			"hashes",
			key+".hc22000",
		)

		err := os.WriteFile(
			filename,
			[]byte(strings.Join(hashes, "\n")+"\n"),
			0644,
		)

		if err != nil {
			return err
		}
	}

	return nil
}
