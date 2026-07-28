package system

import (
	"os"
	"os/exec"
	"path/filepath"
)

type ConversionPlan struct {
	InputType        CaptureFileType
	OutputType       string
	Backend          string
	BackendAvailable bool
	ConversionNeeded bool
	Description      string
	OutputFile       string
	Command          []string
}

func BuildConversionPlan(input string, t CaptureFileType) ConversionPlan {

	_ = os.MkdirAll("output", 0755)

	out := filepath.Base(input)
	out = out[:len(out)-len(filepath.Ext(out))] + ".hc22000"
	out = filepath.Join("output", out)

	switch t {

	case TypePCAP, TypePCAPNG:

		plan := ConversionPlan{
			InputType:        t,
			OutputType:       "HC22000",
			Backend:          "hcxpcapngtool",
			ConversionNeeded: true,
			Description:      "Convert capture to HC22000.",
			OutputFile:       out,
		}

		if _, err := exec.LookPath(plan.Backend); err == nil {
			plan.BackendAvailable = true
			plan.Command = []string{
				plan.Backend,
				"-o", out,
				input,
			}
		}

		return plan

	case TypeHC22000:

		return ConversionPlan{
			InputType:        t,
			OutputType:       "HC22000",
			Backend:          "-",
			BackendAvailable: true,
			ConversionNeeded: false,
			Description:      "Already compatible with Hashcat.",
			OutputFile:       input,
		}
	}

	return ConversionPlan{
		InputType:        TypeUnknown,
		OutputType:       "-",
		Backend:          "-",
		ConversionNeeded: false,
		Description:      "Unsupported format.",
	}
}
