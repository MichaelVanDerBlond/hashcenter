package project

type Network struct {
	ID int

	ESSID string
	BSSID string

	Handshake int
	PMKID     int

	HashFile string
}
