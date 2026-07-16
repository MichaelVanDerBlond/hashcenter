package system

type DiskInfo struct {
	Path      string
	Total     uint64
	Free      uint64
	Available uint64
}
