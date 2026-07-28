package workflow

import "github.com/MichaelVanDerBlond/hashcenter/internal/analyze"

func Build(r *analyze.Report) *Plan {
	p := &Plan{}

	if r == nil {
		p.Add("Analysis report is unavailable.")
		return p
	}

	if r.Conversion {
		p.Add("Convert capture to HC22000 using hcxpcapngtool.")
	}

	if r.HashcatReady {
		p.Add("Capture is ready for Hashcat.")
	}

	if !r.HashcatReady && !r.Conversion {
		p.Add("No action required.")
	}

	return p
}
