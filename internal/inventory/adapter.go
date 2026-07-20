package inventory

type Adapter struct {
	ID int

	Interface string
	Phy       string

	MAC    string
	Driver string

	Mode string

	MonitorSupported bool

	Bands    []string
	Channels []int

	Up bool
}

type Inventory struct {
	Adapters []Adapter
}
