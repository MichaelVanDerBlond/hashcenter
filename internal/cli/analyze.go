package cli

import (
	"fmt"
	"os"

	"github.com/MichaelVanDerBlond/hashcenter/internal/analyze"
)

func Analyze() error {

	var target string

	switch len(os.Args) {

	case 2:

		if _, err := os.Stat("reference/test.cap"); err == nil {
			target = "reference/test.cap"
		} else {
			fmt.Println("Usage:")
			fmt.Println("  hashcenter analyze <capture>")
			fmt.Println()
			fmt.Println("No reference/test.cap found.")
			return nil
		}

	case 3:

		target = os.Args[2]

	default:

		fmt.Println("Usage:")
		fmt.Println("  hashcenter analyze <capture>")
		return nil
	}

	r, err := analyze.Analyze(target)
	if err != nil {
		return err
	}

	fmt.Println("========================================")
	fmt.Println("         HASHCENTER ANALYZE")
	fmt.Println("========================================")
	fmt.Println()

	fmt.Println("[File]")
	fmt.Printf("Path              : %s\n", r.Path)
	fmt.Printf("Size              : %.2f MB\n", float64(r.Size)/(1024*1024))
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

	if r.TShark {
		fmt.Println()
		fmt.Println("[TShark]")
		fmt.Printf("Version           : %s\n", r.Version)
		fmt.Printf("Frames            : %s\n", r.Frames)
		fmt.Printf("Beacon Frames     : %d\n", r.BeaconFrames)
		fmt.Printf("Probe Frames      : %d\n", r.ProbeFrames)
		fmt.Printf("EAPOL Frames      : %d\n", r.EAPOLFrames)
	}

	fmt.Println()

	fmt.Println("[Backends]")

	for _, b := range r.Backends {

		icon := "✖"

		if b.Available {
			icon = "✔"
		}

		fmt.Printf("%s %-18s\n", icon, b.Name)
	}

	fmt.Println()

	fmt.Println("[Next Step]")

	switch {

	case r.HashcatReady:
		fmt.Println("Ready for Hashcat.")

	case r.Conversion:
		fmt.Printf("hashcenter convert %s --execute\n", target)

	default:
		fmt.Println("No recommendation.")
	}

	return nil
}
