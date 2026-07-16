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
		return doctor.Warning("Hashcat", "not installed")
	}

	return doctor.OK("Hashcat", strings.TrimSpace(string(out)))
}
