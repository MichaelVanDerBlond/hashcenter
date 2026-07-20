package inventory

import (
	"bufio"
	"os/exec"
	"strconv"
	"strings"
)

func discover() ([]Adapter, error) {

	cmd := exec.Command("iw", "dev")

	out, err := cmd.Output()
	if err != nil {
		return nil, err
	}

	var adapters []Adapter
	var current *Adapter

	scanner := bufio.NewScanner(strings.NewReader(string(out)))

	for scanner.Scan() {

		line := strings.TrimSpace(scanner.Text())

		switch {

		case strings.HasPrefix(line, "phy#"):

			if current != nil {
				adapters = append(adapters, *current)
			}

			current = &Adapter{}
			current.Phy = strings.TrimPrefix(line, "phy#")

		case strings.HasPrefix(line, "Interface "):

			if current != nil {
				current.Interface = strings.TrimSpace(strings.TrimPrefix(line, "Interface"))
			}

		case strings.HasPrefix(line, "addr "):

			if current != nil {
				current.MAC = strings.TrimSpace(strings.TrimPrefix(line, "addr"))
			}

		case strings.HasPrefix(line, "type "):

			if current != nil {
				current.Mode = strings.TrimSpace(strings.TrimPrefix(line, "type"))
			}
		}
	}

	if current != nil {
		adapters = append(adapters, *current)
	}

	for i := range adapters {
		adapters[i].ID = i + 1

		if _, err := strconv.Atoi(adapters[i].Phy); err == nil {
			adapters[i].Phy = "phy" + adapters[i].Phy
		}
	}

	return adapters, nil
}
