package device

import (
	"fmt"

	"github.com/MichaelVanDerBlond/hashcenter/internal/inventory"
)

type Manager struct {
	Adapter inventory.Adapter
}

func Default() (*Manager, error) {

	inv, err := inventory.List()
	if err != nil {
		return nil, err
	}

	if len(inv.Adapters) == 0 {
		return nil, fmt.Errorf("no wireless adapters found")
	}

	return &Manager{
		Adapter: inv.Adapters[0],
	}, nil
}

func (m *Manager) Interface() string {
	return m.Adapter.Interface
}

func (m *Manager) Driver() string {
	return m.Adapter.Driver
}

func (m *Manager) MonitorSupported() bool {
	return m.Adapter.MonitorSupported
}

func (m *Manager) Mode() string {
	return m.Adapter.Mode
}

func (m *Manager) IsUp() bool {
	return m.Adapter.Up
}
