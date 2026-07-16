package system

import (
	"os"
)

type Analysis struct {
	Path             string
	Size             int64
	Type             CaptureFileType
	HashcatReady     bool
	ConversionNeeded bool
	RecommendedTool  string
}

func AnalyzeFile(path string) (*Analysis, error) {

	stat, err := os.Stat(path)
	if err != nil {
		return nil, err
	}

	t := DetectCaptureFile(path)

	a := &Analysis{
		Path: path,
		Size: stat.Size(),
		Type: t,
	}

	switch t {

	case TypeHC22000:
		a.HashcatReady = true

	case TypePCAP, TypePCAPNG:
		a.ConversionNeeded = true
		a.RecommendedTool = "hcxpcapngtool"

	case TypeHCCAP, TypeHCCAPX:
		a.ConversionNeeded = true
		a.RecommendedTool = "hcxhashtool"
	}

	return a, nil
}
