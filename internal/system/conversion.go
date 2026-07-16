package system

type ConversionPlan struct {
	InputType        CaptureFileType
	OutputType       string
	Backend          string
	ConversionNeeded bool
	Description      string
}

func BuildConversionPlan(t CaptureFileType) ConversionPlan {

	switch t {

	case TypePCAP, TypePCAPNG:
		return ConversionPlan{
			InputType:        t,
			OutputType:       "HC22000",
			Backend:          "hcxpcapngtool",
			ConversionNeeded: true,
			Description:      "Convert capture to HC22000.",
		}

	case TypeHC22000:
		return ConversionPlan{
			InputType:        t,
			OutputType:       "HC22000",
			Backend:          "-",
			ConversionNeeded: false,
			Description:      "Already compatible with Hashcat.",
		}

	case Type16800:
		return ConversionPlan{
			InputType:        t,
			OutputType:       "16800",
			Backend:          "-",
			ConversionNeeded: false,
			Description:      "PMKID hash detected.",
		}

	case Type16801:
		return ConversionPlan{
			InputType:        t,
			OutputType:       "16801",
			Backend:          "-",
			ConversionNeeded: false,
			Description:      "PMKID + ESSID hash detected.",
		}

	case TypeHCCAP:
		return ConversionPlan{
			InputType:        t,
			OutputType:       "HC22000",
			Backend:          "hcxhashtool",
			ConversionNeeded: true,
			Description:      "Legacy HCCAP conversion required.",
		}

	case TypeHCCAPX:
		return ConversionPlan{
			InputType:        t,
			OutputType:       "HC22000",
			Backend:          "hcxhashtool",
			ConversionNeeded: true,
			Description:      "Legacy HCCAPX conversion required.",
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
