package cli

import (
	"fmt"

	"github.com/MichaelVanDerBlond/hashcenter/internal/project"
)

func List() error {

	list, err := project.Load()
	if err != nil {
		return err
	}

	fmt.Println("========================================")
	fmt.Println("          HASHCENTER PROJECT")
	fmt.Println("========================================")
	fmt.Println()

	if len(list) == 0 {

		fmt.Println("No converted networks found.")
		fmt.Println("Run:")
		fmt.Println("  hashcenter convert <capture>")

		return nil
	}

	fmt.Printf("%-3s %-32s %-17s %-3s %-5s\n",
		"#",
		"ESSID",
		"BSSID",
		"HS",
		"PMKID",
	)

	fmt.Println("----------------------------------------------------------------------------")

	for _, n := range list {

		fmt.Printf(
			"%-3d %-32s %-17s %-3d %-5d\n",
			n.ID,
			n.ESSID,
			n.BSSID,
			n.Handshake,
			n.PMKID,
		)
	}

	return nil
}
