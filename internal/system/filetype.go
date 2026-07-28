package system

import (
	"bufio"
	"bytes"
	"io"
	"os"
	"path/filepath"
	"strings"
)

type CaptureFileType string

const (
	TypeUnknown CaptureFileType = "UNKNOWN"
	TypePCAP    CaptureFileType = "PCAP"
	TypePCAPNG  CaptureFileType = "PCAPNG"
	TypeHC22000 CaptureFileType = "HC22000"
	TypeHCCAPX  CaptureFileType = "HCCAPX"
	TypeHCCAP   CaptureFileType = "HCCAP"
	Type16800   CaptureFileType = "16800"
	Type16801   CaptureFileType = "16801"
)

func DetectCaptureFile(path string) CaptureFileType {

	f, err := os.Open(path)
	if err == nil {
		defer f.Close()

		header := make([]byte, 32)
		n, _ := f.Read(header)
		header = header[:n]

		switch {

		case bytes.HasPrefix(header, []byte{0xd4, 0xc3, 0xb2, 0xa1}),
			bytes.HasPrefix(header, []byte{0xa1, 0xb2, 0xc3, 0xd4}),
			bytes.HasPrefix(header, []byte{0x4d, 0x3c, 0xb2, 0xa1}),
			bytes.HasPrefix(header, []byte{0xa1, 0xb2, 0x3c, 0x4d}):

			return TypePCAP

		case bytes.HasPrefix(header, []byte{0x0A, 0x0D, 0x0D, 0x0A}):

			return TypePCAPNG
		}

		if _, err := f.Seek(0, io.SeekStart); err == nil {

			scanner := bufio.NewScanner(f)

			if scanner.Scan() {

				line := scanner.Text()

				switch {

				case strings.HasPrefix(line, "WPA*"):
					return TypeHC22000

				case strings.Contains(line, ":16800:"):
					return Type16800

				case strings.Contains(line, ":16801:"):
					return Type16801
				}
			}
		}
	}

	switch strings.ToLower(filepath.Ext(path)) {

	case ".cap", ".pcap":
		return TypePCAP

	case ".pcapng":
		return TypePCAPNG

	case ".hc22000", ".22000":
		return TypeHC22000

	case ".16800":
		return Type16800

	case ".16801":
		return Type16801

	case ".hccap":
		return TypeHCCAP

	case ".hccapx":
		return TypeHCCAPX
	}

	return TypeUnknown
}
