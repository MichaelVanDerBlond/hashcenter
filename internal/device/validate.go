package device

import "fmt"

func (m *Manager) Validate() error {

	if m.Adapter.Interface == "" {
		return fmt.Errorf("wireless interface not found")
	}

	if !m.Adapter.Up {
		return fmt.Errorf(
			"interface %s is down",
			m.Adapter.Interface,
		)
	}

	if !m.Adapter.MonitorSupported {
		return fmt.Errorf(
			"adapter %s does not support monitor mode",
			m.Adapter.Interface,
		)
	}

	return nil
}
