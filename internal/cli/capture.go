package cli

import (
	"fmt"

	"github.com/MichaelVanDerBlond/hashcenter/internal/system"
)

func Capture() error {

	info, err := system.Collect()
	if err != nil {
		return err
	}

	fmt.Println("========================================")
	fmt.Println("        HASHCENTER CAPTURE")
	fmt.Println("========================================")
	fmt.Println()

	if len(info.WiFi) == 0 {
		fmt.Println("No Wi-Fi interfaces detected.")
		return nil
	}

	fmt.Printf("%-4s %-10s %-20s %-12s\n",
		"#",
		"PHY",
		"INTERFACE",
		"MODE")

	fmt.Println("--------------------------------------------------------")

	for i, iface := range info.WiFi {
		fmt.Printf("%-4d %-10s %-20s %-12s\n",
			i+1,
			iface.Phy,
			iface.Name,
			iface.Type,
		)
	}

	return nil
}
