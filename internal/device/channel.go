package device

import (
	"fmt"
	"strconv"
	"strings"
)

func (m *Manager) CurrentChannel() (string, error) {

	out, err := run("iw", "dev", m.Adapter.Interface, "info")
	if err != nil {
		return "", err
	}

	for _, line := range strings.Split(out, "\n") {

		line = strings.TrimSpace(line)

		if strings.HasPrefix(line, "channel ") {

			fields := strings.Fields(line)
			if len(fields) >= 2 {
				return fields[1], nil
			}
		}
	}

	return "", fmt.Errorf("channel not detected")
}

func (m *Manager) SetChannel(channel int) error {

	if channel < 1 || channel > 196 {
		return fmt.Errorf("invalid channel: %d", channel)
	}

	if _, err := run(
		"iw",
		"dev",
		m.Adapter.Interface,
		"set",
		"channel",
		strconv.Itoa(channel),
	); err != nil {
		return err
	}

	return nil
}
