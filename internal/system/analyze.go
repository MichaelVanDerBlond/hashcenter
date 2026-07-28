package system

import (
	"os"
	"os/exec"
	"strings"
)

type Analysis struct {
	Path             string
	Size             int64
	Type             CaptureFileType
	HashcatReady     bool
	ConversionNeeded bool
	RecommendedTool  string

	CapinfosAvailable bool
	PacketCount       string
	CaptureDuration   string
	FileEncapsulation string
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

	if t == TypePCAP || t == TypePCAPNG {

		if _, err := exec.LookPath("capinfos"); err == nil {

			a.CapinfosAvailable = true

			out, err := exec.Command("capinfos", path).CombinedOutput()
			if err == nil {

				value := func(line, prefix string) string {
					return strings.TrimSpace(strings.TrimPrefix(line, prefix))
				}

				for _, line := range strings.Split(string(out), "\n") {

					switch {

					case strings.HasPrefix(line, "Number of packets:"):
						a.PacketCount = value(line, "Number of packets:")

					case strings.HasPrefix(line, "Capture duration:"):
						a.CaptureDuration = value(line, "Capture duration:")

					case strings.HasPrefix(line, "File encapsulation:"):
						a.FileEncapsulation = value(line, "File encapsulation:")
					}
				}
			}
		}
	}

	return a, nil
}
