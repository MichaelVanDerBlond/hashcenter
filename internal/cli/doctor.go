package cli

import (
	"fmt"

	"github.com/MichaelVanDerBlond/hashcenter/internal/doctor"
	"github.com/MichaelVanDerBlond/hashcenter/internal/doctor/checks"
)

func Doctor() error {
	d := doctor.Doctor{
		Checks: []doctor.Check{
			checks.OS{},
			checks.Kernel{},
			checks.Hashcat{},
		},
	}

	results := d.Run()

	for _, r := range results {
		status := "✖"
		if r.Passed {
			status = "✔"
		}

		fmt.Printf("%s %-20s %s\n", status, r.Name, r.Message)
	}

	return nil
}
