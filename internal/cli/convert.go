package cli

import (
	"fmt"
	"os"
	"strings"

	"github.com/MichaelVanDerBlond/hashcenter/internal/system"
)

func Convert() error {

	execute := false

	args := os.Args[2:]

	if len(args) == 2 && args[1] == "--execute" {
		execute = true
		args = args[:1]
	}

	if len(args) != 1 {
		fmt.Println("Usage:")
		fmt.Println("  hashcenter convert <file>")
		fmt.Println("  hashcenter convert <file> --execute")
		return nil
	}

	input := args[0]

	stat, err := os.Stat(input)
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
	fmt.Printf("Size               : %.2f KB\n", float64(stat.Size())/1024)
	fmt.Printf("Detected Type      : %s\n", plan.InputType)
	fmt.Printf("Target Type        : %s\n", plan.OutputType)
	fmt.Printf("Backend            : %s\n", plan.Backend)
	fmt.Printf("Backend Available  : %t\n", plan.BackendAvailable)
	fmt.Printf("Conversion Needed  : %t\n", plan.ConversionNeeded)
	fmt.Printf("Output File        : %s\n", plan.OutputFile)

	if len(plan.Command) > 0 {
		fmt.Println()
		fmt.Println("Command:")
		fmt.Printf("  %s\n", strings.Join(plan.Command, " "))
	}

	if !execute {
		fmt.Println()
		fmt.Println("Dry-run mode.")
		fmt.Println("Use --execute to run the conversion.")
		return nil
	}

	if !plan.BackendAvailable {
		return fmt.Errorf("backend %q not found", plan.Backend)
	}

	if len(plan.Command) == 0 {
		fmt.Println()
		fmt.Println("Nothing to execute.")
		return nil
	}

	fmt.Println()
	fmt.Println("Executing...")

	result, err := system.RunCommand(plan.Command)

	fmt.Println(result.Output)

	if err != nil {
		return err
	}

	fmt.Println("Done.")

	return nil
}
