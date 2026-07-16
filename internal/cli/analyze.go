package cli

import (
	"fmt"
	"os"

	"github.com/MichaelVanDerBlond/hashcenter/internal/analyze"
	"github.com/MichaelVanDerBlond/hashcenter/internal/workflow"
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

	plan := workflow.Build(r)

	fmt.Println("========================================")
	fmt.Println("         HASHCENTER ANALYZE")
	fmt.Println("========================================")
	fmt.Println()

	fmt.Println("[File]")
	fmt.Printf("Path              : %s\n", r.Path)
	fmt.Printf("Detected Type     : %s\n", r.Type)
	fmt.Printf("Size              : %.2f MB\n", float64(r.Size)/(1024*1024))

	fmt.Println()

	fmt.Println("[Workflow]")

	for i, step := range plan.Steps {
		fmt.Printf("%d. %s\n", i+1, step)
	}

	fmt.Println()

	fmt.Println("[Capture]")
	fmt.Printf("Packets           : %s\n", r.Packets)
	fmt.Printf("Frames            : %s\n", r.Frames)
	fmt.Printf("Duration          : %s\n", r.Duration)

	fmt.Println()

	fmt.Println("[Wi-Fi]")
	fmt.Printf("Beacon            : %d\n", r.BeaconFrames)
	fmt.Printf("Probe             : %d\n", r.ProbeFrames)
	fmt.Printf("EAPOL             : %d\n", r.EAPOLFrames)

	return nil
}
