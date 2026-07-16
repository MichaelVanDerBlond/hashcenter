package cli

import (
	"fmt"

	"github.com/MichaelVanDerBlond/hashcenter/internal/system"
)

func Doctor() error {
	info, err := system.Collect()
	if err != nil {
		return err
	}

	fmt.Printf("%-2s %-20s %s\n", "✔", "Operating System", info.OS)
	fmt.Printf("%-2s %-20s %s\n", "✔", "Kernel", info.Kernel)

	if info.Hashcat == "not installed" {
		fmt.Printf("%-2s %-20s %s\n", "!", "Hashcat", info.Hashcat)
	} else {
		fmt.Printf("%-2s %-20s %s\n", "✔", "Hashcat", info.Hashcat)
	}

	fmt.Printf("%-2s %-20s %s\n", "✔", "CPU", info.CPU.Model)
	fmt.Printf("%-2s %-20s %d cores\n", "✔", "CPU Cores", info.CPU.Cores)
	fmt.Printf("%-2s %-20s %.2f GB\n",
		"✔",
		"Memory",
		float64(info.Memory.Total)/(1024*1024*1024),
	)

	fmt.Printf("%-2s %-20s %.2f / %.2f GB free\n",
		"✔",
		"Disk",
		float64(info.Disk.Available)/(1024*1024*1024),
		float64(info.Disk.Total)/(1024*1024*1024),
	)
	if info.GPU.Name != "" {
		fmt.Printf("%-2s %-20s %s\n", "✔", "GPU", info.GPU.Name)
		fmt.Printf("%-2s %-20s %s\n", "✔", "Driver", info.GPU.Driver)
		fmt.Printf("%-2s %-20s %s / %s MiB\n",
			"✔",
			"GPU Memory",
			info.GPU.MemoryUsed,
			info.GPU.MemoryTotal,
		)
		fmt.Printf("%-2s %-20s %s °C\n", "✔", "GPU Temp", info.GPU.Temperature)
		fmt.Printf("%-2s %-20s %s %%\n", "✔", "GPU Load", info.GPU.Utilization)
	} else {
		fmt.Printf("%-2s %-20s %s\n", "!", "GPU", "not detected")
	}

	return nil
}
