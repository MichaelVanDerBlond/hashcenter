package checks

import (
	"os/exec"
	"strings"

	"github.com/MichaelVanDerBlond/hashcenter/internal/doctor"
)

type Kernel struct{}

func (Kernel) Name() string {
	return "Kernel"
}

func (Kernel) Run() doctor.Result {
	out, err := exec.Command("uname", "-r").Output()
	if err != nil {
		return doctor.Error("Kernel", err.Error())
	}

	return doctor.OK("Kernel", strings.TrimSpace(string(out)))
}
