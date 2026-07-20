package inventory

import (
	"bufio"
	"os/exec"
	"strings"
)

func enrichIW(adapters []Adapter) ([]Adapter, error) {

	out, err := exec.Command("iw", "phy").Output()
	if err != nil {
		return adapters, err
	}

	scanner := bufio.NewScanner(strings.NewReader(string(out)))

	var phy string
	var inModes bool
	var inBand bool

	for scanner.Scan() {

		line := scanner.Text()
		trim := strings.TrimSpace(line)

		if strings.HasPrefix(trim, "Wiphy ") {
			phy = strings.TrimPrefix(trim, "Wiphy ")
			inModes = false
			inBand = false
			continue
		}

		if trim == "Supported interface modes:" {
			inModes = true
			inBand = false
			continue
		}

		if strings.HasPrefix(trim, "Band ") {
			inBand = true
			inModes = false
			continue
		}

		if inModes && strings.Contains(trim, "* monitor") {
			for i := range adapters {
				if adapters[i].Phy == phy {
					adapters[i].MonitorSupported = true
				}
			}
		}

		if inBand && strings.Contains(trim, "2412.0 MHz") {
			for i := range adapters {
				if adapters[i].Phy == phy && !contains(adapters[i].Bands, "2.4 GHz") {
					adapters[i].Bands = append(adapters[i].Bands, "2.4 GHz")
				}
			}
		}

		if inBand && strings.Contains(trim, "5180.0 MHz") {
			for i := range adapters {
				if adapters[i].Phy == phy && !contains(adapters[i].Bands, "5 GHz") {
					adapters[i].Bands = append(adapters[i].Bands, "5 GHz")
				}
			}
		}
	}

	return adapters, nil
}

func contains(slice []string, value string) bool {
	for _, v := range slice {
		if v == value {
			return true
		}
	}
	return false
}
