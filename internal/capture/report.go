package capture

type Report struct {
	FileType string
	Frames   uint64

	HasBeacon bool
	HasProbe  bool
	HasEAPOL  bool
	HasPMKID  bool
}
