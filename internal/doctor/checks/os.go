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
		return doctor.Error("Operating System", err.Error())
	}

	for _, line := range strings.Split(string(data), "\n") {
		if strings.HasPrefix(line, "PRETTY_NAME=") {
			value := strings.Trim(strings.TrimPrefix(line, "PRETTY_NAME="), "\"")
			return doctor.OK("Operating System", value)
		}
	}

	return doctor.Error("Operating System", "PRETTY_NAME not found")
}
