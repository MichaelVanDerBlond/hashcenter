package device

func (m *Manager) EnableMonitor() error {

	if err := m.Validate(); err != nil {
		return err
	}

	if m.Adapter.Mode == "monitor" {
		return nil
	}

	if _, err := run(
		"ip",
		"link",
		"set",
		m.Adapter.Interface,
		"down",
	); err != nil {
		return err
	}

	if _, err := run(
		"iw",
		m.Adapter.Interface,
		"set",
		"monitor",
		"none",
	); err != nil {
		return err
	}

	if _, err := run(
		"ip",
		"link",
		"set",
		m.Adapter.Interface,
		"up",
	); err != nil {
		return err
	}

	m.Adapter.Mode = "monitor"
	m.Adapter.Up = true

	return nil
}

func (m *Manager) DisableMonitor() error {

	if m.Adapter.Mode != "monitor" {
		return nil
	}

	if _, err := run(
		"ip",
		"link",
		"set",
		m.Adapter.Interface,
		"down",
	); err != nil {
		return err
	}

	if _, err := run(
		"iw",
		m.Adapter.Interface,
		"set",
		"type",
		"managed",
	); err != nil {
		return err
	}

	if _, err := run(
		"ip",
		"link",
		"set",
		m.Adapter.Interface,
		"up",
	); err != nil {
		return err
	}

	m.Adapter.Mode = "managed"
	m.Adapter.Up = true

	return nil
}
