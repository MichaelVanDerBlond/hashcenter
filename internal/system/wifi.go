package system

import (
	"os/exec"
	"strings"
)

type WiFiInterface struct {
	Name string
}

func collectWiFi(info *Info) error {

	out, err := exec.Command("iw", "dev").Output()
	if err != nil {
		return nil
	}

	lines := strings.Split(string(out), "\n")

	for _, line := range lines {

		line = strings.TrimSpace(line)

		if strings.HasPrefix(line, "Interface ") {

			info.WiFi = append(info.WiFi, WiFiInterface{
				Name: strings.TrimSpace(strings.TrimPrefix(line, "Interface")),
			})
		}
	}

	return nil
}
