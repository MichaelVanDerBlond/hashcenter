package cli

import (
	"fmt"
	"os"
	"strings"

	"github.com/MichaelVanDerBlond/hashcenter/internal/system"
)

func Convert() error {

	if len(os.Args) != 3 {
		fmt.Println("Usage:")
		fmt.Println("  hashcenter convert <file>")
		return nil
	}

	input := os.Args[2]

	info, err := os.Stat(input)
	if err != nil {
		return err
	}

	fileType := system.DetectCaptureFile(input)
	plan := system.BuildConversionPlan(input, fileType)

	fmt.Println("========================================")
	fmt.Println("         HASHCENTER CONVERT")
	fmt.Println("========================================")
	fmt.Println()

	fmt.Printf("Input File         : %s\n", input)
	fmt.Printf("Size               : %.2f KB\n", float64(info.Size())/1024)
	fmt.Printf("Detected Type      : %s\n", plan.InputType)
	fmt.Printf("Target Type        : %s\n", plan.OutputType)
	fmt.Printf("Backend            : %s\n", plan.Backend)

	if plan.BackendAvailable {
		fmt.Println("Backend Status     : AVAILABLE")
	} else {
		fmt.Println("Backend Status     : NOT FOUND")
	}

	fmt.Printf("Conversion Needed  : %t\n", plan.ConversionNeeded)
	fmt.Printf("Output File        : %s\n", plan.OutputFile)
	fmt.Printf("Description        : %s\n", plan.Description)

	if len(plan.Command) > 0 {
		fmt.Println()
		fmt.Println("Planned command:")
		fmt.Printf("  %s\n", strings.Join(plan.Command, " "))
	}

	return nil
}
