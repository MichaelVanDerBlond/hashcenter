package device

type Status struct {
	Interface string
	Driver    string
	Mode      string
	Up        bool
	Monitor   bool
}

func (m *Manager) Status() Status {

	return Status{
		Interface: m.Adapter.Interface,
		Driver:    m.Adapter.Driver,
		Mode:      m.Adapter.Mode,
		Up:        m.Adapter.Up,
		Monitor:   m.Adapter.MonitorSupported,
	}
}
