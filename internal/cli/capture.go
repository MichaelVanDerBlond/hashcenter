package cli

import (
	"fmt"
	"os"

	"github.com/MichaelVanDerBlond/hashcenter/internal/system"
)

func Capture() error {
	info, err := system.Collect()
	if err != nil {
		return err
	}

	// hashcenter capture
	if len(os.Args) == 2 {
		printInterfaces(info)
		return nil
	}

	// hashcenter capture start <interface>
	if len(os.Args) == 4 && os.Args[2] == "start" {
		return captureStart(info, os.Args[3])
	}

	fmt.Println("Usage:")
	fmt.Println("  hashcenter capture")
	fmt.Println("  hashcenter capture start <interface>")
	return nil
}

func printInterfaces(info *system.Info) {
	fmt.Println("========================================")
	fmt.Println("        HASHCENTER CAPTURE")
	fmt.Println("========================================")
	fmt.Println()

	if len(info.WiFi) == 0 {
		fmt.Println("No Wi-Fi interfaces detected.")
		return
	}

	fmt.Printf("%-3s %-8s %-18s %-10s\n",
		"#", "PHY", "INTERFACE", "MODE")
	fmt.Println("------------------------------------------------")

	for i, iface := range info.WiFi {
		fmt.Printf("%-3d %-8s %-18s %-10s\n",
			i+1,
			iface.Phy,
			iface.Name,
			iface.Type,
		)
	}
}

func captureStart(info *system.Info, name string) error {
	var iface *system.WiFiInterface

	for i := range info.WiFi {
		if info.WiFi[i].Name == name {
			iface = &info.WiFi[i]
			break
		}
	}

	if iface == nil {
		return fmt.Errorf("interface %q not found", name)
	}

	fmt.Println("========================================")
	fmt.Println("      HASHCENTER CAPTURE (DRY RUN)")
	fmt.Println("========================================")
	fmt.Println()

	fmt.Printf("Interface : %s\n", iface.Name)
	fmt.Printf("PHY       : %s\n", iface.Phy)
	fmt.Printf("Mode      : %s\n", iface.Type)
	fmt.Println()

	fmt.Println("Planned actions:")
	fmt.Println("  [1] Verify interface state")
	fmt.Println("  [2] Verify monitor-mode capability")
	fmt.Println("  [3] Prepare capture backend")
	fmt.Println()
	fmt.Println("Dry-run only: no system changes were made.")

	return nil
}
