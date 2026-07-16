package analyze

import "github.com/MichaelVanDerBlond/hashcenter/internal/network"

type Backend struct {
	Name      string
	Available bool
}

type Report struct {
	Path string
	Size int64

	Type string

	HashcatReady bool
	Conversion   bool

	Backend string

	Backends []Backend

	Networks network.List

	Capinfos bool

	Packets       string
	Duration      string
	Encapsulation string

	TShark  bool
	Version string

	Frames string

	BeaconFrames uint64
	ProbeFrames  uint64
	EAPOLFrames  uint64
}
