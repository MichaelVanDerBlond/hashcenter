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
		return doctor.Result{
			Name:    "Kernel",
			Passed:  false,
			Message: err.Error(),
		}
	}

	return doctor.Result{
		Name:    "Kernel",
		Passed:  true,
		Message: strings.TrimSpace(string(out)),
	}
}
