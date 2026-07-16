package system

import (
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

	ext := strings.ToLower(filepath.Ext(path))

	switch ext {

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

	case ".hccapx":
		return TypeHCCAPX

	case ".hccap":
		return TypeHCCAP

	default:
		return TypeUnknown
	}
}
