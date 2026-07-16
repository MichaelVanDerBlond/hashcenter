package analyze

type Report struct {
	Path string
	Size int64

	Type string

	HashcatReady bool
	Conversion   bool

	Backend string

	Capinfos bool

	Packets       string
	Duration      string
	Encapsulation string
}
