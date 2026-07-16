package cli

import (
	"fmt"
	"os"

	"github.com/MichaelVanDerBlond/hashcenter/internal/system"
)

func Convert() error {

	if len(os.Args) != 3 {
		fmt.Println("Usage:")
		fmt.Println("  hashcenter convert <file>")
		return nil
	}

	path := os.Args[2]

	info, err := os.Stat(path)
	if err != nil {
		return err
	}

	fileType := system.DetectCaptureFile(path)

	fmt.Println("========================================")
	fmt.Println("         HASHCENTER CONVERT")
	fmt.Println("========================================")
	fmt.Println()

	fmt.Printf("File : %s\n", path)
	fmt.Printf("Size : %.2f KB\n", float64(info.Size())/1024)
	fmt.Printf("Type : %s\n\n", fileType)

	switch fileType {

	case system.TypePCAP:
		fmt.Println("Capture format detected.")
		fmt.Println("Recommended output : HC22000")
		fmt.Println("Backend            : hcxpcapngtool")

	case system.TypePCAPNG:
		fmt.Println("Capture format detected.")
		fmt.Println("Recommended output : HC22000")
		fmt.Println("Backend            : hcxpcapngtool")

	case system.TypeHC22000:
		fmt.Println("Already ready for Hashcat.")

	case system.Type16800:
		fmt.Println("PMKID hash detected.")

	case system.Type16801:
		fmt.Println("PMKID + ESSID hash detected.")

	case system.TypeHCCAP:
		fmt.Println("Legacy HCCAP detected.")

	case system.TypeHCCAPX:
		fmt.Println("Legacy HCCAPX detected.")

	default:
		fmt.Println("Unsupported or unknown format.")
	}

	return nil
}
