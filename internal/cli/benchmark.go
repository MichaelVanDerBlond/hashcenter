package cli

import (
	"fmt"

	"github.com/MichaelVanDerBlond/hashcenter/internal/system"
)

func Benchmark() error {
	info, err := system.Collect()
	if err != nil {
		return err
	}

	fmt.Println("========================================")
	fmt.Println("           HASHCENTER BENCHMARK")
	fmt.Println("========================================")
	fmt.Println()

	score := 0
	maxScore := 5

	// GPU
	fmt.Print("GPU          : ")
	if info.GPU.Name != "" {
		fmt.Printf("OK (%s)\n", info.GPU.Name)
		score++
	} else {
		fmt.Println("FAIL")
	}

	// Driver
	fmt.Print("Driver       : ")
	if info.GPU.Driver != "" {
		fmt.Printf("OK (%s)\n", info.GPU.Driver)
		score++
	} else {
		fmt.Println("FAIL")
	}

	// Memory
	fmt.Print("Memory       : ")
	memGB := float64(info.Memory.Total) / (1024 * 1024 * 1024)
	if memGB >= 8 {
		fmt.Printf("OK (%.2f GB)\n", memGB)
		score++
	} else {
		fmt.Printf("LOW (%.2f GB)\n", memGB)
	}

	// CPU
	fmt.Print("CPU Threads  : ")
	if info.CPU.Threads >= 4 {
		fmt.Printf("OK (%d)\n", info.CPU.Threads)
		score++
	} else {
		fmt.Printf("LOW (%d)\n", info.CPU.Threads)
	}

	// Hashcat
	fmt.Print("Hashcat      : ")

	hashcatOK := false
	for _, tool := range info.Tools {
		if tool.Name == "Hashcat" && tool.Present {
			fmt.Printf("OK (%s)\n", tool.Version)
			hashcatOK = true
			score++
			break
		}
	}

	if !hashcatOK {
		fmt.Println("NOT INSTALLED")
	}

	fmt.Println()
	fmt.Printf("Score : %d/%d\n", score, maxScore)

	switch {
	case score == maxScore:
		fmt.Println("Status: READY")
	case score >= 3:
		fmt.Println("Status: USABLE")
	default:
		fmt.Println("Status: NOT READY")
	}

	return nil
}
