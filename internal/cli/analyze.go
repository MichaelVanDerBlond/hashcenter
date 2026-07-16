package cli

import (
	"fmt"
	"os"

	"github.com/MichaelVanDerBlond/hashcenter/internal/system"
)

func Analyze() error {

	if len(os.Args) != 3 {
		fmt.Println("Usage:")
		fmt.Println("  hashcenter analyze <file>")
		return nil
	}

	result, err := system.AnalyzeFile(os.Args[2])
	if err != nil {
		return err
	}

	fmt.Println("========================================")
	fmt.Println("         HASHCENTER ANALYZE")
	fmt.Println("========================================")
	fmt.Println()

	fmt.Println("[File]")
	fmt.Printf("Path              : %s\n", result.Path)
	fmt.Printf("Size              : %.2f KB\n", float64(result.Size)/1024)
	fmt.Printf("Detected Type     : %s\n", result.Type)
	fmt.Println()

	fmt.Println("[Analysis]")
	fmt.Printf("Hashcat Ready     : %t\n", result.HashcatReady)
	fmt.Printf("Conversion Needed : %t\n", result.ConversionNeeded)

	if result.RecommendedTool != "" {
		fmt.Printf("Recommended Tool  : %s\n", result.RecommendedTool)
	}

	fmt.Println()

	fmt.Println("[Next Step]")

	switch {
	case result.HashcatReady:
		fmt.Println("Use the file directly with Hashcat.")

	case result.ConversionNeeded:
		fmt.Printf("hashcenter convert %s --execute\n", result.Path)

	default:
		fmt.Println("No recommendation available.")
	}

	return nil
}
