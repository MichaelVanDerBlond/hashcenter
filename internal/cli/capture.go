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

	fmt.Println("Detected Wi-Fi interfaces")
	fmt.Println()

	for i, iface := range info.WiFi {

		fmt.Printf("%d. %s\n", i+1, iface.Name)
	}

	return nil
}
