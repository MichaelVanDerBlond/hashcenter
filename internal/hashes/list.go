package hashes

import (
	"encoding/hex"
	"os"
	"path/filepath"
	"sort"
	"strings"
)

type HashFile struct {
	BSSID string
	ESSID string

	Handshakes int
	PMKIDs     int

	Path string
}

func List(dir string) ([]HashFile, error) {

	entries, err := os.ReadDir(dir)
	if err != nil {
		return nil, err
	}

	var out []HashFile

	for _, entry := range entries {

		if entry.IsDir() {
			continue
		}

		if filepath.Ext(entry.Name()) != ".hc22000" {
			continue
		}

		path := filepath.Join(dir, entry.Name())

		data, err := os.ReadFile(path)
		if err != nil {
			continue
		}

		var h HashFile

		h.Path = path

		lines := strings.Split(strings.TrimSpace(string(data)), "\n")

		for i, line := range lines {

			line = strings.TrimSpace(line)

			if line == "" {
				continue
			}

			fields := strings.Split(line, "*")

			if len(fields) < 6 {
				continue
			}

			if i == 0 {

				h.BSSID = strings.ToUpper(fields[3])

				h.ESSID = "<hidden>"

				if b, err := hex.DecodeString(fields[5]); err == nil {

					if s := strings.TrimSpace(string(b)); s != "" {
						h.ESSID = s
					}
				}
			}

			switch fields[1] {

			case "01":
				h.PMKIDs++

			case "02":
				h.Handshakes++

			}
		}

		out = append(out, h)
	}

	sort.Slice(out, func(i, j int) bool {

		if out[i].Handshakes != out[j].Handshakes {
			return out[i].Handshakes > out[j].Handshakes
		}

		if out[i].PMKIDs != out[j].PMKIDs {
			return out[i].PMKIDs > out[j].PMKIDs
		}

		return out[i].ESSID < out[j].ESSID
	})

	return out, nil
}
