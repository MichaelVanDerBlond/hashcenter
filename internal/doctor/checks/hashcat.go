package checks

import (
	"os/exec"
	"strings"

	"github.com/MichaelVanDerBlond/hashcenter/internal/doctor"
)

type Hashcat struct{}

func (Hashcat) Name() string {
	return "Hashcat"
}

func (Hashcat) Run() doctor.Result {
	out, err := exec.Command("hashcat", "--version").Output()
	if err != nil {
		return doctor.Result{
			Name:    "Hashcat",
			Passed:  false,
			Message: "not installed",
		}
	}

	return doctor.Result{
		Name:    "Hashcat",
		Passed:  true,
		Message: strings.TrimSpace(string(out)),
	}
}
