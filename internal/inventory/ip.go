package inventory

import (
	"bufio"
	"os/exec"
	"strings"
)

func enrichIP(adapters []Adapter) ([]Adapter, error) {

	out, err := exec.Command("ip", "link").Output()
	if err != nil {
		return adapters, err
	}

	scanner := bufio.NewScanner(strings.NewReader(string(out)))

	for scanner.Scan() {

		line := scanner.Text()

		if strings.HasPrefix(line, " ") {
			continue
		}

		fields := strings.SplitN(line, ":", 3)
		if len(fields) < 3 {
			continue
		}

		iface := strings.TrimSpace(fields[1])

		state := strings.Contains(line, "UP")

		for i := range adapters {
			if adapters[i].Interface == iface {
				adapters[i].Up = state
				break
			}
		}
	}

	return adapters, nil
}
