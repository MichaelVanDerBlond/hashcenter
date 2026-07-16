package system

import "syscall"

func collectDisk(info *Info) error {
	var stat syscall.Statfs_t

	if err := syscall.Statfs("/", &stat); err != nil {
		return err
	}

	bsize := uint64(stat.Bsize)

	info.Disk = DiskInfo{
		Path:      "/",
		Total:     stat.Blocks * bsize,
		Free:      stat.Bfree * bsize,
		Available: stat.Bavail * bsize,
	}

	return nil
}
