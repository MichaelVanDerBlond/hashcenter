package cli

import (
	"fmt"
	"strings"

	"github.com/MichaelVanDerBlond/hashcenter/internal/inventory"
)

func Inventory() error {

	inv, err := inventory.List()
	if err != nil {
		return err
	}

	fmt.Println("========================================")
	fmt.Println("        HASHCENTER INVENTORY")
	fmt.Println("========================================")
	fmt.Println()

	fmt.Printf("%-3s %-18s %-6s %-12s %-10s %-6s %-8s %s\n",
		"ID",
		"INTERFACE",
		"PHY",
		"DRIVER",
		"MODE",
		"STATE",
		"MONITOR",
		"BANDS",
	)

	fmt.Println("----------------------------------------------------------------------------")

	for _, a := range inv.Adapters {

		state := "DOWN"
		if a.Up {
			state = "UP"
		}

		monitor := "NO"
		if a.MonitorSupported {
			monitor = "YES"
		}

		bands := "-"
		if len(a.Bands) > 0 {
			bands = strings.Join(a.Bands, ", ")
		}

		fmt.Printf(
			"%-3d %-18s %-6s %-12s %-10s %-6s %-8s %s\n",
			a.ID,
			a.Interface,
			a.Phy,
			a.Driver,
			a.Mode,
			state,
			monitor,
			bands,
		)
	}

	return nil
}
