package cli

import (
	"fmt"

	"github.com/MichaelVanDerBlond/hashcenter/internal/system"
)

func Inventory() error {
	info, err := system.Collect()
	if err != nil {
		return err
	}

	fmt.Println("========================================")
	fmt.Println("           HASHCENTER INVENTORY")
	fmt.Println("========================================")
	fmt.Println()

	fmt.Println("[System]")
	fmt.Printf("OS           : %s\n", info.OS)
	fmt.Printf("Kernel       : %s\n", info.Kernel)
	fmt.Println()

	fmt.Println("[CPU]")
	fmt.Printf("Vendor       : %s\n", info.CPU.Vendor)
	fmt.Printf("Model        : %s\n", info.CPU.Model)
	fmt.Printf("Architecture : %s\n", info.CPU.Arch)
	fmt.Printf("Cores        : %d\n", info.CPU.Cores)
	fmt.Printf("Threads      : %d\n", info.CPU.Threads)
	fmt.Println()

	fmt.Println("[Memory]")
	fmt.Printf("Total        : %.2f GB\n", float64(info.Memory.Total)/(1024*1024*1024))
	fmt.Printf("Available    : %.2f GB\n", float64(info.Memory.Available)/(1024*1024*1024))
	fmt.Printf("Free         : %.2f GB\n", float64(info.Memory.Free)/(1024*1024*1024))
	fmt.Println()

	fmt.Println("[Disk]")
	fmt.Printf("Mount        : %s\n", info.Disk.Path)
	fmt.Printf("Total        : %.2f GB\n", float64(info.Disk.Total)/(1024*1024*1024))
	fmt.Printf("Available    : %.2f GB\n", float64(info.Disk.Available)/(1024*1024*1024))
	fmt.Println()

	fmt.Println("[GPU]")
	if info.GPU.Name == "" {
		fmt.Println("Not detected")
	} else {
		fmt.Printf("Model        : %s\n", info.GPU.Name)
		fmt.Printf("Driver       : %s\n", info.GPU.Driver)
		fmt.Printf("Memory       : %s / %s MiB\n", info.GPU.MemoryUsed, info.GPU.MemoryTotal)
		fmt.Printf("Temperature  : %s °C\n", info.GPU.Temperature)
		fmt.Printf("Load         : %s %%\n", info.GPU.Utilization)
	}
	fmt.Println()

	fmt.Println("[Tools]")
	for _, tool := range info.Tools {
		if tool.Present {
			fmt.Printf("[OK]   %-15s %s\n", tool.Name, tool.Version)
		} else {
			fmt.Printf("[MISS] %-15s not installed\n", tool.Name)
		}
	}

	return nil
}
