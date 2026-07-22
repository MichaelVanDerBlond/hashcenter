package cli

import (
	"fmt"
	"os"
)

func Execute() error {

	if len(os.Args) < 2 {
		usage()
		return nil
	}

	switch os.Args[1] {

	case "analyze":
		return Analyze()

	case "convert":
		return Convert()

	case "list":
		return List()

	case "inventory":
		return Inventory()

	case "capture":
		return Capture()

	case "attack":
		return Attack()

	case "sessions":
		return Sessions()

	case "status":
		return Status()

	case "doctor":
		return Doctor()

	case "benchmark":
		return Benchmark()

	case "version":
		return Version()

	case "serve":
		return Serve()

	default:
		usage()
		return nil
	}
}

func usageAnalyze() {
	fmt.Println("Usage:")
	fmt.Println("  hashcenter analyze <capture>")
}

func usageConvert() {
	fmt.Println("Usage:")
	fmt.Println("  hashcenter convert <capture>")
}

func usage() {

	fmt.Println("HashCenter")
	fmt.Println()
	fmt.Println("Usage:")
	fmt.Println("  hashcenter analyze <capture>")
	fmt.Println("  hashcenter convert <capture>")
	fmt.Println("  hashcenter list")
	fmt.Println("  hashcenter inventory")
	fmt.Println("  hashcenter capture --channel 6 --time 30 --output office")
	fmt.Println("  hashcenter attack")
	fmt.Println("  hashcenter sessions")
	fmt.Println("  hashcenter status")
	fmt.Println("  hashcenter doctor")
	fmt.Println("  hashcenter benchmark")
	fmt.Println("  hashcenter version")
	fmt.Println("  hashcenter serve")
}
