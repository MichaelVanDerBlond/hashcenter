package analyze

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

	Capinfos bool

	Packets       string
	Duration      string
	Encapsulation string

	TShark  bool
	Version string

	Frames string
}
