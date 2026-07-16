package hc22000

type Network struct {
	BSSID string
	ESSID string

	HasPMKID bool
	HasEAPOL bool
}
