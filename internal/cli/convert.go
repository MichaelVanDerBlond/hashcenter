package cli

import (
	"fmt"
	"os"

	"github.com/MichaelVanDerBlond/hashcenter/internal/hashes"
	"github.com/MichaelVanDerBlond/hashcenter/internal/splitter"
	"github.com/MichaelVanDerBlond/hashcenter/internal/system"
)

func Convert() error {

	if len(os.Args) != 3 {
		usageConvert()
		return nil
	}

	input := os.Args[2]

	stat, err := os.Stat(input)
	if err != nil {
		return err
	}

	fmt.Println("========================================")
	fmt.Println("         HASHCENTER CONVERT")
	fmt.Println("========================================")
	fmt.Println()

	fmt.Printf("Input File        : %s\n", input)
	fmt.Printf("Size              : %.2f MB\n", float64(stat.Size())/(1024*1024))
	fmt.Println()

	fmt.Println("Converting...")

	plan, _, err := system.ExecuteConversion(input)
	if err != nil {
		return err
	}

	fmt.Println("Done.")
	fmt.Printf("HC22000           : %s\n", plan.OutputFile)

	fmt.Println()
	fmt.Println("Splitting...")

	if err := splitter.Split(plan.OutputFile); err != nil {
		return err
	}

	list, err := hashes.List("output/hashes")
	if err != nil {
		return err
	}

	fmt.Println("Done.")
	fmt.Println()

	fmt.Printf("[Ready Networks] (%d)\n", len(list))
	fmt.Printf("%-3s %-32s %-3s %-5s\n", "#", "ESSID", "HS", "PMKID")
	fmt.Println("-----------------------------------------------------------")

	for i, h := range list {
		fmt.Printf(
			"%-3d %-32s %-3d %-5d\n",
			i+1,
			h.ESSID,
			h.Handshakes,
			h.PMKIDs,
		)
	}

	fmt.Println()
	fmt.Println("Hashes directory : output/hashes")

	return nil
}
