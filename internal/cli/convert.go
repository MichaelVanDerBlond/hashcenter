package cli

import (
	"fmt"
	"os"

	"github.com/MichaelVanDerBlond/hashcenter/internal/system"
)

func Convert() error {

	if len(os.Args) != 3 {
		fmt.Println("Usage:")
		fmt.Println("  hashcenter convert <file>")
		return nil
	}

	path := os.Args[2]

	stat, err := os.Stat(path)
	if err != nil {
		return err
	}

	fileType := system.DetectCaptureFile(path)
	plan := system.BuildConversionPlan(fileType)

	fmt.Println("========================================")
	fmt.Println("         HASHCENTER CONVERT")
	fmt.Println("========================================")
	fmt.Println()

	fmt.Printf("Input File         : %s\n", path)
	fmt.Printf("Size               : %.2f KB\n", float64(stat.Size())/1024)
	fmt.Printf("Detected Type      : %s\n", plan.InputType)
	fmt.Printf("Target Type        : %s\n", plan.OutputType)
	fmt.Printf("Backend            : %s\n", plan.Backend)
	fmt.Printf("Conversion Needed  : %t\n", plan.ConversionNeeded)
	fmt.Printf("Description        : %s\n", plan.Description)

	return nil
}
