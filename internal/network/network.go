package network

type Network struct {
	ID int

	BSSID string
	ESSID string

	Security string

	HandshakeCount int

	HasPMKID bool

	Ready bool
}
