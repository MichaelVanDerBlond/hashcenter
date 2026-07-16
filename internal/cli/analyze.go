package cli

import (
	"fmt"
	"os"

	"github.com/MichaelVanDerBlond/hashcenter/internal/analyze"
)

func Analyze() error {

	if len(os.Args) != 3 {
		fmt.Println("Usage:")
		fmt.Println("  hashcenter analyze <file>")
		return nil
	}

	r, err := analyze.Analyze(os.Args[2])
	if err != nil {
		return err
	}

	fmt.Println("========================================")
	fmt.Println("         HASHCENTER ANALYZE")
	fmt.Println("========================================")
	fmt.Println()

	fmt.Println("[File]")
	fmt.Printf("Path              : %s\n", r.Path)
	fmt.Printf("Size              : %.2f KB\n", float64(r.Size)/1024)
	fmt.Printf("Detected Type     : %s\n", r.Type)
	fmt.Println()

	fmt.Println("[Analysis]")
	fmt.Printf("Hashcat Ready     : %t\n", r.HashcatReady)
	fmt.Printf("Conversion Needed : %t\n", r.Conversion)

	if r.Backend != "" {
		fmt.Printf("Backend           : %s\n", r.Backend)
	}

	if r.Capinfos {

		fmt.Println()
		fmt.Println("[Capture]")
		fmt.Printf("Packets           : %s\n", r.Packets)
		fmt.Printf("Duration          : %s\n", r.Duration)
		fmt.Printf("Encapsulation     : %s\n", r.Encapsulation)
	}

	fmt.Println()
	fmt.Println("[Next Step]")

	if r.HashcatReady {
		fmt.Println("Ready for Hashcat.")
	} else if r.Conversion {
		fmt.Printf("hashcenter convert %s --execute\n", r.Path)
	} else {
		fmt.Println("No recommendation.")
	}

	return nil
}
