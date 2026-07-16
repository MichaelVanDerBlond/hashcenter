package cli

import (
	"fmt"
	"os"
)

func Execute() error {
	if len(os.Args) < 2 {
		printHelp()
		return nil
	}

	switch os.Args[1] {

	case "version":
		return Version()

	case "doctor":
		return Doctor()

	case "inventory":
		return Inventory()

	case "capture":
		return Capture()

	case "convert":
		return Convert()

	case "analyze":
		return Analyze()

	case "benchmark":
		return Benchmark()

	case "help":
		printHelp()
		return nil

	default:
		fmt.Printf("Unknown command: %s\n\n", os.Args[1])
		printHelp()
		return fmt.Errorf("unknown command")
	}
}

func printHelp() {
	fmt.Println("HashCenter")
	fmt.Println()
	fmt.Println("Usage:")
	fmt.Println("  hashcenter <command>")
	fmt.Println()
	fmt.Println("Commands:")
	fmt.Println("  version")
	fmt.Println("  doctor")
	fmt.Println("  inventory")
	fmt.Println("  capture")
	fmt.Println("  convert")
	fmt.Println("  analyze")
	fmt.Println("  benchmark")
}
