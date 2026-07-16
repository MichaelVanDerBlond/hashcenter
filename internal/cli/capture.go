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

	fmt.Printf("%-3s %-8s %-18s %-10s %-9s %-8s\n",
		"#",
		"PHY",
		"INTERFACE",
		"MODE",
		"MONITOR",
		"AIRMON")

	fmt.Println("----------------------------------------------------------------")

	for i, iface := range info.WiFi {

		monitor := "YES"
		if !iface.MonitorCapable {
			monitor = "NO"
		}

		airmon := "YES"
		if !iface.AirmonAvailable {
			airmon = "NO"
		}

		fmt.Printf("%-3d %-8s %-18s %-10s %-9s %-8s\n",
			i+1,
			iface.Phy,
			iface.Name,
			iface.Type,
			monitor,
			airmon,
		)
	}

	fmt.Println()
	fmt.Println("Next step:")
	fmt.Println("  hashcenter capture start <interface>")

	return nil
}
