package capture

import (
	"encoding/binary"
	"fmt"
	"io"
	"os"
)

const (
	magicPCAPLE = 0xa1b2c3d4
	magicPCAPBE = 0xd4c3b2a1
	magicPCAPNG = 0x0A0D0D0A
)

type Reader struct {
	file *os.File
}

func Open(path string) (*Reader, *Report, error) {

	f, err := os.Open(path)
	if err != nil {
		return nil, nil, err
	}

	var magic uint32

	if err := binary.Read(f, binary.LittleEndian, &magic); err != nil {
		f.Close()
		return nil, nil, err
	}

	rep := &Report{}

	switch magic {

	case magicPCAPLE, magicPCAPBE:
		rep.FileType = "PCAP"

	case magicPCAPNG:
		rep.FileType = "PCAPNG"

	default:
		f.Close()
		return nil, nil, fmt.Errorf("unsupported capture format")
	}

	if _, err := f.Seek(0, io.SeekStart); err != nil {
		f.Close()
		return nil, nil, err
	}

	return &Reader{file: f}, rep, nil
}

func (r *Reader) Close() error {
	if r == nil || r.file == nil {
		return nil
	}
	return r.file.Close()
}
