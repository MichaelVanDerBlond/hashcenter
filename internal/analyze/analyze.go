package analyze

import (
	"os"

	"github.com/MichaelVanDerBlond/hashcenter/internal/network"
	"github.com/MichaelVanDerBlond/hashcenter/internal/system"
)

func Analyze(path string) (*Report, error) {

	stat, err := os.Stat(path)
	if err != nil {
		return nil, err
	}

	fileType := system.DetectCaptureFile(path)

	r := &Report{
		Path: path,
		Size: stat.Size(),
		Type: string(fileType),
	}

	switch fileType {

	case system.TypeHC22000:
		r.HashcatReady = true

	case system.TypePCAP, system.TypePCAPNG:
		r.Conversion = true
		r.Backend = "hcxpcapngtool"

	case system.TypeHCCAP, system.TypeHCCAPX:
		r.Conversion = true
		r.Backend = "hcxhashtool"
	}

	_ = collectCapinfos(r)
	_ = collectTShark(r)
	_ = collectHCX(r)

	if nets, err := network.Discover(path); err == nil {
		r.Networks = *nets
	}

	return r, nil
}
