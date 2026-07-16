package checks

import (
	"os"
	"strings"

	"github.com/MichaelVanDerBlond/hashcenter/internal/doctor"
)

type OS struct{}

func (OS) Name() string {
	return "Operating System"
}

func (OS) Run() doctor.Result {
	data, err := os.ReadFile("/etc/os-release")
	if err != nil {
		return doctor.Result{
			Name:    "Operating System",
			Passed:  false,
			Message: err.Error(),
		}
	}

	for _, line := range strings.Split(string(data), "\n") {
		if strings.HasPrefix(line, "PRETTY_NAME=") {
			value := strings.TrimPrefix(line, "PRETTY_NAME=")
			value = strings.Trim(value, "\"")

			return doctor.Result{
				Name:    "Operating System",
				Passed:  true,
				Message: value,
			}
		}
	}

	return doctor.Result{
		Name:    "Operating System",
		Passed:  false,
		Message: "PRETTY_NAME not found",
	}
}
