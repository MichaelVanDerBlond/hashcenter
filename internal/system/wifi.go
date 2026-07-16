package system

import (
	"os/exec"
	"strings"
)

type WiFiInterface struct {
	Name string
	Phy  string
	Type string
}

func collectWiFi(info *Info) error {

	out, err := exec.Command("iw", "dev").Output()
	if err != nil {
		return nil
	}

	var current *WiFiInterface

	for _, line := range strings.Split(string(out), "\n") {

		line = strings.TrimSpace(line)

		switch {

		case strings.HasPrefix(line, "phy#"):
			info.WiFi = append(info.WiFi, WiFiInterface{
				Phy: line,
			})
			current = &info.WiFi[len(info.WiFi)-1]

		case strings.HasPrefix(line, "Interface ") && current != nil:
			current.Name = strings.TrimSpace(strings.TrimPrefix(line, "Interface"))

		case strings.HasPrefix(line, "type ") && current != nil:
			current.Type = strings.TrimSpace(strings.TrimPrefix(line, "type"))
		}
	}

	return nil
}
