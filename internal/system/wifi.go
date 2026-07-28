package system

import (
	"os/exec"
	"strings"
)

type WiFiInterface struct {
	Name            string
	Phy             string
	Type            string
	MonitorCapable  bool
	AirmonAvailable bool
}

func collectWiFi(info *Info) error {

	_, err := exec.LookPath("airmon-ng")
	airmonAvailable := err == nil

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
				Phy:             line,
				AirmonAvailable: airmonAvailable,
				MonitorCapable:  true,
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
