package network

import (
	"bufio"
	"bytes"
	"os/exec"
	"strings"

	"github.com/MichaelVanDerBlond/hashcenter/internal/hashes"
)

func Discover(path string) (*List, error) {

	list := &List{}

	out, err := exec.Command(
		"tshark",
		"-r", path,
		"-Y", "wlan.fc.type_subtype==8",
		"-T", "fields",
		"-e", "wlan.bssid",
		"-e", "wlan.ssid",
	).CombinedOutput()

	if err != nil {
		return list, err
	}

	ready := make(map[string]hashes.HashFile)

	if files, err := hashes.List("output/hashes"); err == nil {
		for _, h := range files {
			ready[strings.ToUpper(h.BSSID)] = h
		}
	}

	seen := make(map[string]bool)

	scanner := bufio.NewScanner(bytes.NewReader(out))

	for scanner.Scan() {

		fields := strings.Split(scanner.Text(), "\t")

		if len(fields) < 2 {
			continue
		}

		bssid := strings.ToUpper(strings.TrimSpace(fields[0]))
		essid := DecodeESSID(fields[1])

		if bssid == "" {
			continue
		}

		if seen[bssid] {
			continue
		}

		seen[bssid] = true

		n := Network{
			BSSID: bssid,
			ESSID: essid,
		}

		if h, ok := ready[bssid]; ok {
			n.Ready = true
			n.HandshakeCount = h.Handshakes
			n.HasPMKID = h.PMKIDs > 0
		}

		list.Add(n)
	}

	if err := scanner.Err(); err != nil {
		return list, err
	}

	return list, nil
}
